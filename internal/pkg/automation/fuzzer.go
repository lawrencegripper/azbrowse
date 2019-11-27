package automation

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/config"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
	"github.com/stuartleeks/gocui"
)

// ErrorReporter Tracks errors during fuzzing
type ErrorReporter struct {
}

// Errorf adds an error found while fuzzing
func (e *ErrorReporter) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(5)
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

	testFunc := func(currentNode *expanders.TreeNode, nodes []*expanders.TreeNode) {
		// if r := regexp.MustCompile(".*/providers/Microsoft.ContainerRegistry/registries/.*/"); r.MatchString(currentNode.ID) {
		// 	st.Expect(&ErrorReporter{}, len(nodes), 5)
		// }
	}

	shouldSkip := func(currentNode *expanders.TreeNode, itemId string) (shouldSkip bool, maxToExpand int) {
		///
		/// Limit walking of things that have LOTS of nodes
		/// so we don't get stuck
		///

		// Only expand limitted set under container repositories
		if r := regexp.MustCompile(".*/<repositories>/.*/.*"); r.MatchString(itemId) {
			return false, 1
		}

		// Only expand limitted set under activity log
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/<activitylog>"); r.MatchString(itemId) {
			return false, 1
		}

		// Only expand limitted set under deployments
		if r := regexp.MustCompile("/subscriptions/.*/resourceGroups/.*/providers/Microsoft.Resources/deployments"); r.MatchString(itemId) {
			return false, 1
		}

		// Only expand limitted set under metrics
		if strings.HasSuffix(itemId, "providers/microsoft.insights/metricdefinitions") {
			return false, 1
		}

		return false, -1
	}

	stateMap := map[string]*fuzzState{}

	startTime := time.Now()
	endTime := startTime.Add(time.Duration(settings.FuzzerDurationMinutes) * time.Minute)
	go func() {
		navigatedChannel := eventing.SubscribeToTopic("list.navigated")

		// If used with `-navigate` wait for navigation to finish before fuzzing
		if settings.NavigateToID != "" {
			for {
				if !processNavigations {
					<-time.After(time.Second * 2)
					list.ExpandCurrentSelection()
					break
				}
			}
		}

		stack := []*views.ListNavigatedEventState{}

		for {
			if time.Now().After(endTime) {
				gui.Close()
				os.Exit(0)
			}

			navigateStateInterface := <-navigatedChannel
			<-time.After(200 * time.Millisecond)

			navigateState := navigateStateInterface.(views.ListNavigatedEventState)
			nodeList := navigateState.NewNodes

			stack = append(stack, &navigateState)

			state, exists := stateMap[navigateState.ParentNodeID]
			if !exists {
				state = &fuzzState{
					current: 0,
					max:     len(nodeList),
				}
				stateMap[navigateState.ParentNodeID] = state
			}

			if state.current >= state.max {
				// Pop stack
				if len(stack) > 1 {
					stack = stack[:len(stack)-1]
				}

				// Navigate back
				list.GoBack()
				continue
			}

			currentItem := list.GetNodes()[state.current]
			skipItem, maxItem := shouldSkip(currentItem, navigateState.ParentNodeID)
			if maxItem != -1 {
				state.max = min(maxItem, state.max)
			}

			var resultNodes []*expanders.TreeNode
			if skipItem {
				// Skip the current item and expand
				state.current++

				// If skip takes us to the end of the available items nav back up stack
				if state.current >= state.max {
					// Pop stack
					stack = stack[:len(stack)-1]

					// Navigate back
					list.GoBack()
					continue
				}

				list.ChangeSelection(state.current)
				list.ExpandCurrentSelection()

				resultNodes = list.GetNodes()
			} else {
				list.ChangeSelection(state.current)
				list.ExpandCurrentSelection()
				state.current++

				resultNodes = list.GetNodes()
			}

			testFunc(currentItem, resultNodes)
		}
	}()
}
