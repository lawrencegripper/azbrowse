package automation

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/nbio/st"
	"github.com/stuartleeks/gocui"
)

// ErrorReporter Tracks errors during fuzzing
type ErrorReporter struct {
	name string
}

func testName(n string) *ErrorReporter {
	return &ErrorReporter{name: n}
}

// Errorf adds an error found while fuzzing
func (e *ErrorReporter) Errorf(format string, args ...interface{}) {
	intro := fmt.Sprintf("Fuzzer Test `%s` Failed: ", e.name)
	log.Panicf(intro+format, args...)
}

func assertNavigationCorrect(expandedItem *expanders.TreeNode, resultingNodes []*expanders.TreeNode) {

	itemIDSegmentLength := len(strings.Split(expandedItem.ID, "/"))

	// Assert container registry handled correctly
	if r := regexp.MustCompile(".*/Microsoft.ContainerRegistry/registries/.*"); r.MatchString(expandedItem.ID) && itemIDSegmentLength == 9 {
		expectedNodesInACR := 13
		st.Expect(testName("containerRegistry_root_assertExpandNodeCount"), len(resultingNodes), expectedNodesInACR)
	}

	// Add more tests here...
}

// StartAutomatedFuzzer will start walking the tree of nodes
// logging out information and running tests while walking the tree
func StartAutomatedFuzzer(list *views.ListWidget, settings *config.Settings, gui *gocui.Gui) {
	shouldSkipMetricsProvider := false
	shouldSkip := func(currentNode *expanders.TreeNode, itemID string) (shouldSkip bool) {
		///
		/// Limit walking of things that have LOTS of nodes
		/// so we don't get stuck
		///

		itemIDSegmentLength := len(strings.Split(currentNode.ID, "/"))

		// Don't walk too deeply into any tree
		if itemIDSegmentLength > 11 {
			return true
		}

		// Skip walking all processes on every instance of a webapp!
		if r := regexp.MustCompile(".*/sites/.*/processes"); r.MatchString(currentNode.ID) {
			return true
		}

		// Skip container repos
		if r := regexp.MustCompile(".*/<repositories>/.*/.*"); r.MatchString(itemID) {
			return true
		}

		// Skip activity log
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/<activitylog>"); r.MatchString(itemID) {
			return true
		}

		// Skip expanding individual deployments
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/providers/Microsoft.Resources/deployments/.*"); r.MatchString(itemID) && itemIDSegmentLength >= 7 {
			return true
		}

		// Only expand limitted set under metrics
		if strings.HasSuffix(itemID, "providers/microsoft.insights/metricdefinitions") {
			defer func() {
				shouldSkipMetricsProvider = true
			}()
			return shouldSkipMetricsProvider
		}

		return false
	}

	visitedNodes := map[string]*expanders.TreeNode{}

	startTime := time.Now()
	endTime := startTime.Add(time.Duration(settings.FuzzerDurationMinutes) * time.Minute)
	go func() {
		// recover from panic, if one occurrs, and leave terminal usable
		defer errorhandling.RecoveryWithCleanup()

		var navigatedChannel chan interface{}

		// If used with `-navigate` wait for navigation to finish before fuzzing
		if settings.NavigateToID != "" {
			for {
				<-time.After(time.Second * 1)
				if !navigateToInProgress {

					// `-navigate` is finished, subscribe to nav events and get started
					// by expanding the current item
					navigatedChannel = eventing.SubscribeToTopic("list.navigated")
					list.ExpandCurrentSelection()

					break
				}
			}
		} else {
			navigatedChannel = eventing.SubscribeToTopic("list.navigated")
		}

		for {
			if time.Now().After(endTime) {
				gui.Close()
				fmt.Println("Fuzzer completed with no panics")
				os.Exit(0)
			}

			navigateStateInterface := <-navigatedChannel

			navigateState := navigateStateInterface.(views.ListNavigatedEventState)

			// If started with `-navigate` don't walk outside the specified resource
			if navigateState.ParentNodeID != "root" && settings.NavigateToID != "" && !strings.HasPrefix(navigateState.ParentNodeID, settings.NavigateToID) {
				fmt.Println("Fuzzer finished working on nodes under `-navigate` ID supplied")
				st.Expect(testName("EXPECTED ERROR limit_fuzz_to_navigate_node_id_fuzz_completed"), navigateState.ParentNodeID, settings.NavigateToID)
			}

			currentNodes := list.GetNodes()
			newNodeFound := false
			for index, node := range currentNodes {
				_, alreadyVisited := visitedNodes[node.ID]
				if !alreadyVisited {
					skip := shouldSkip(node, navigateState.ParentNodeID)
					if skip {
						visitedNodes[node.ID] = node
						continue
					}

					list.ChangeSelection(index)
					// Expand it
					list.ExpandCurrentSelection()
					visitedNodes[node.ID] = node
					newNodeFound = true

					// Assert things about the navigation based on the current item and the result nodes
					assertNavigationCorrect(node, list.GetNodes())
					break
				}
			}

			// All nodes in the list already visited or skipped
			if !newNodeFound {
				list.GoBack()
			}

		}
	}()
}
