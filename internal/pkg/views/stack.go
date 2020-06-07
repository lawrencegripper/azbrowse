package views

import "github.com/lawrencegripper/azbrowse/internal/pkg/expanders"

// Page represents a previous view in the nav stack
type Page struct {
	Value            []*expanders.TreeNode
	Data             string
	DataType         expanders.ExpanderResponseType
	Title            string
	Selection        int
	FilterString     string
	ExpandedNodeItem *expanders.TreeNode
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Page
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Page) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Page {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
