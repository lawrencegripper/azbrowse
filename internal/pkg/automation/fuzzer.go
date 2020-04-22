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
		st.Expect(testName("containerRegistry_root_assertExpandHas5Nodes"), len(resultingNodes), expectedNodesInACR)
	}

	// Add more tests here...
}

// StartAutomatedFuzzer will start walking the tree of nodes
// logging out information and running tests while walking the tree
func StartAutomatedFuzzer(list *views.ListWidget, settings *config.Settings, gui *gocui.Gui) {

	type fuzzState struct {
		current int
		max     int
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	shouldSkipActivityProvider := false
	shouldSkipMetricsProvider := false
	shouldSkip := func(currentNode *expanders.TreeNode, itemID string) (shouldSkip bool, maxToExpand int) {
		///
		/// Limit walking of things that have LOTS of nodes
		/// so we don't get stuck
		///

		itemIDSegmentLength := len(strings.Split(currentNode.ID, "/"))

		// Only expand limitted set under container repositories
		if r := regexp.MustCompile(".*/<repositories>/.*/.*"); r.MatchString(itemID) {
			return false, 1
		}

		// Only expand limitted set under activity log
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/<activitylog>"); r.MatchString(itemID) {
			// Only walk the activity provider the first time we see it.
			defer func() {
				shouldSkipActivityProvider = true
			}()
			return shouldSkipActivityProvider, 1
		}

		// Only expand limitted set under deployments
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/providers/Microsoft.Resources/deployments"); r.MatchString(itemID) && itemIDSegmentLength >= 7 {
			return false, 1
		}

		// Only expand limitted set under metrics
		if strings.HasSuffix(itemID, "providers/microsoft.insights/metricdefinitions") {
			defer func() {
				shouldSkipMetricsProvider = true
			}()
			return shouldSkipMetricsProvider, 1
		}

		return false, -1
	}

	stateMap := map[string]*fuzzState{}

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
				if !processNavigations {

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
			if settings.NavigateToID != "" && !strings.HasPrefix(navigateState.ParentNodeID, settings.NavigateToID) {
				fmt.Println("Fuzzer finished working on nodes under `-navigate` ID supplied")
				st.Expect(testName("EXPECTED_limit_fuzz_to_navigate_node_id"), navigateState.ParentNodeID, settings.NavigateToID)
			}

			nodeList := navigateState.NewNodes

			// Create or Retrieve the current status of the fuzzer for this
			// level in the tree
			state, exists := stateMap[navigateState.ParentNodeID]
			if !exists {
				state = &fuzzState{
					current: 0,
					max:     len(nodeList),
				}
				stateMap[navigateState.ParentNodeID] = state
			}

			// Fuzzing completed on this level
			if state.current >= state.max {
				// Navigate back
				list.GoBack()
				continue
			}

			// Get the current item to decide if we should skip or limit how many
			// children to expand
			currentItem := list.GetNodes()[state.current]
			skipItem, maxItem := shouldSkip(currentItem, navigateState.ParentNodeID)
			if maxItem != -1 {
				state.max = min(maxItem, state.max)
			}

			// Store the resulting nodes so we can assert tests about expandedItem -> resultingNodes
			// behaved as expected
			var resultNodes []*expanders.TreeNode

			if skipItem {
				// Skip the current item and expand
				state.current++

				// If skip takes us to the end of the available items nav back up stack
				if state.current >= state.max {
					// Navigate back
					list.GoBack()
					continue
				}

				// Move to next item
				list.ChangeSelection(state.current)

				// Expand it
				list.ExpandCurrentSelection()

				resultNodes = list.GetNodes()
			} else {
				// Move to current item
				list.ChangeSelection(state.current)
				// Expand it
				list.ExpandCurrentSelection()
				state.current++

				resultNodes = list.GetNodes()
			}

			// Assert things about the navigation based on the current item and the result nodes
			assertNavigationCorrect(currentItem, resultNodes)
		}
	}()
}
