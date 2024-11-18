package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

// createTestScreen creates a new simulation screen for testing
func createTestScreen(t *testing.T) tcell.Screen {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatalf("failed to initialize screen: %v", err)
	}
	return screen
}

// createSampleTree creates a sample tree for testing
func createSampleTree() *TreeNode {
	root := &TreeNode{Text: "Root"}
	child1 := &TreeNode{Text: "Child1"}
	child2 := &TreeNode{Text: "Child2"}
	grandchild1 := &TreeNode{Text: "Grandchild1"}
	grandchild2 := &TreeNode{Text: "Grandchild2"}

	// Set up parent-child relationships
	root.Children = []*TreeNode{child1, child2}
	child1.Parent = root
	child2.Parent = root

	child1.Children = []*TreeNode{grandchild1, grandchild2}
	grandchild1.Parent = child1
	grandchild2.Parent = child1

	return root
}

// TestTreeViewInitialization tests the initialization of TreeView
func TestTreeViewInitialization(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	if tv == nil {
		t.Fatal("NewTreeView returned nil")
	}

	// Check default values
	if !tv.ShowLines {
		t.Error("ShowLines should be true by default")
	}
	if tv.Indent != 2 {
		t.Error("Default indent should be 2")
	}
	if tv.Root != nil {
		t.Error("Root should be nil initially")
	}
	if tv.Selected != nil {
		t.Error("Selected should be nil initially")
	}
}

// TestTreeViewNodeOperations tests node operations
func TestTreeViewNodeOperations(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)

	// Test SetRoot
	root := &TreeNode{Text: "Root"}
	tv.SetRoot(root)
	if tv.Root != root {
		t.Error("SetRoot failed to set the root node")
	}
	if tv.Selected != nil {
		t.Error("SetRoot should not set Selected")
	}

	// Test AddNode
	child := tv.AddNode(root, "Child")
	if child == nil {
		t.Fatal("AddNode returned nil")
	}
	if len(root.Children) != 1 || root.Children[0] != child {
		t.Error("AddNode failed to add child to parent")
	}
	if child.Parent != root {
		t.Error("AddNode failed to set child's parent")
	}
	if child.Text != "Child" {
		t.Error("AddNode failed to set child's text")
	}

	// Test RemoveNode
	tv.RemoveNode(child)
	if len(root.Children) != 0 {
		t.Error("RemoveNode failed to remove child from parent")
	}
	if child.Parent != nil {
		t.Error("RemoveNode failed to clear child's parent")
	}
}

// TestTreeViewExpansion tests expansion and collapse operations
func TestTreeViewExpansion(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	root := createSampleTree()
	tv.SetRoot(root)

	// Test ExpandAll
	tv.ExpandAll()
	var checkExpanded func(*TreeNode) bool
	checkExpanded = func(node *TreeNode) bool {
		if !node.Expanded {
			return false
		}
		for _, child := range node.Children {
			if !checkExpanded(child) {
				return false
			}
		}
		return true
	}
	if !checkExpanded(root) {
		t.Error("ExpandAll failed to expand all nodes")
	}

	// Test CollapseAll
	tv.CollapseAll()
	var checkCollapsed func(*TreeNode) bool
	checkCollapsed = func(node *TreeNode) bool {
		if node.Expanded {
			return false
		}
		for _, child := range node.Children {
			if !checkCollapsed(child) {
				return false
			}
		}
		return true
	}
	if !checkCollapsed(root) {
		t.Error("CollapseAll failed to collapse all nodes")
	}

	// Test ToggleSelected
	tv.Selected = root
	if !tv.ToggleSelected() {
		t.Error("ToggleSelected should return true when node is selected")
	}
	if !root.Expanded {
		t.Error("ToggleSelected failed to expand selected node")
	}
	if !tv.ToggleSelected() {
		t.Error("ToggleSelected should return true when node is selected")
	}
	if root.Expanded {
		t.Error("ToggleSelected failed to collapse selected node")
	}
}

// TestTreeViewNavigation tests navigation operations
func TestTreeViewNavigation(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	root := createSampleTree()
	tv.SetRoot(root)
	root.Expanded = true

	// Test initial selection
	if !tv.SelectNext() {
		t.Error("Initial SelectNext should select root")
	}
	if tv.Selected != root {
		t.Error("Initial SelectNext should select root node")
	}

	// Test navigation with expanded nodes
	child1 := root.Children[0]
	child2 := root.Children[1]

	if !tv.SelectNext() {
		t.Error("SelectNext should select first child when root is expanded")
	}
	if tv.Selected != child1 {
		t.Error("SelectNext should select first child")
	}

	if !tv.SelectNext() {
		t.Error("SelectNext should move to next sibling")
	}
	if tv.Selected != child2 {
		t.Error("SelectNext should select second child")
	}

	// Test SelectPrevious
	if !tv.SelectPrevious() {
		t.Error("SelectPrevious should return true from second child")
	}
	if tv.Selected != child1 {
		t.Error("SelectPrevious should select previous sibling")
	}

	if !tv.SelectPrevious() {
		t.Error("SelectPrevious should return true from first child")
	}
	if tv.Selected != root {
		t.Error("SelectPrevious should select parent")
	}

	// Test deep navigation
	child1.Expanded = true
	grandchild := child1.Children[0]

	tv.Selected = child1
	if !tv.SelectNext() {
		t.Error("SelectNext should work with expanded child")
	}
	if tv.Selected != grandchild {
		t.Error("SelectNext should select grandchild")
	}

	// Test navigation at boundaries
	if tv.SelectNext() {
		t.Error("SelectNext should return false at leaf node")
	}
	if tv.Selected != grandchild {
		t.Error("SelectNext should not change selection at leaf node")
	}

	if !tv.SelectPrevious() {
		t.Error("SelectPrevious should work from leaf node")
	}
	if tv.Selected != child1 {
		t.Error("SelectPrevious should select parent from leaf")
	}
}

// TestTreeViewKeyboardEvents tests keyboard event handling
func TestTreeViewKeyboardEvents(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	root := createSampleTree()
	tv.SetRoot(root)
	root.Expanded = true // Ensure root is expanded initially

	// Test up/down navigation
	tv.Selected = root.Children[0] // Start with first child selected
	ev := tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	handled := tv.HandleKeyEvent(ev)
	if !handled {
		t.Error("Up arrow should be handled")
	}
	if tv.Selected != root {
		t.Error("Up arrow failed to select parent")
	}

	ev = tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	handled = tv.HandleKeyEvent(ev)
	if !handled {
		t.Error("Down arrow should be handled")
	}
	if tv.Selected != root.Children[0] {
		t.Error("Down arrow failed to select first child")
	}

	// Test expand/collapse
	tv.Selected = root
	root.Expanded = true
	ev = tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone)
	handled = tv.HandleKeyEvent(ev)
	if !handled {
		t.Error("Left arrow should be handled")
	}
	if root.Expanded {
		t.Error("Left arrow failed to collapse root")
	}

	ev = tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone)
	handled = tv.HandleKeyEvent(ev)
	if !handled {
		t.Error("Right arrow should be handled")
	}
	if !root.Expanded {
		t.Error("Right arrow failed to expand root")
	}
}

// TestTreeViewScrolling tests scrolling functionality
func TestTreeViewScrolling(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	// Create a tree view with small height to test scrolling
	tv := NewTreeView(screen, 0, 0, 40, 3) // Height of 3 to ensure scrolling is needed
	root := createSampleTree()
	tv.SetRoot(root)
	root.Expanded = true
	root.Children[0].Expanded = true

	// Test initial scroll position
	if tv.ScrollOffset != 0 {
		t.Error("Initial scroll offset should be 0")
	}

	// Test ScrollTo
	tv.ScrollTo(2)
	if tv.ScrollOffset != 2 {
		t.Error("ScrollTo failed to set correct offset")
	}

	// Test ScrollBy
	tv.ScrollBy(-1)
	if tv.ScrollOffset != 1 {
		t.Error("ScrollBy failed to adjust offset")
	}

	// Test ScrollTo with negative value
	tv.ScrollTo(-1)
	if tv.ScrollOffset != 0 {
		t.Error("ScrollTo should clamp negative values to 0")
	}

	// Test EnsureVisible with node outside view
	tv.Selected = root.Children[0].Children[0] // Select a grandchild
	tv.ScrollTo(0)                            // Scroll to top
	tv.EnsureVisible()
	
	// Calculate expected position
	expectedPos := 2 // Root(0) -> Child1(1) -> Grandchild1(2)
	expectedOffset := expectedPos - tv.Height + 1
	if expectedOffset < 0 {
		expectedOffset = 0
	}
	
	if tv.ScrollOffset != expectedOffset {
		t.Errorf("EnsureVisible set wrong scroll offset: got %d, want %d", tv.ScrollOffset, expectedOffset)
	}
}

// TestTreeViewSearch tests node search functionality
func TestTreeViewSearch(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	root := createSampleTree()
	tv.SetRoot(root)

	// Test FindNode
	node := tv.FindNode("Grandchild1")
	if node == nil {
		t.Error("FindNode failed to find existing node")
	}
	if node.Text != "Grandchild1" {
		t.Error("FindNode returned wrong node")
	}

	// Test FindNode with non-existent node
	node = tv.FindNode("NonExistent")
	if node != nil {
		t.Error("FindNode should return nil for non-existent node")
	}

	// Test FindAndSelect
	if !tv.FindAndSelect("Grandchild2") {
		t.Error("FindAndSelect should return true for existing node")
	}
	if tv.Selected == nil || tv.Selected.Text != "Grandchild2" {
		t.Error("FindAndSelect failed to select correct node")
	}
	if !root.Children[0].Expanded {
		t.Error("FindAndSelect failed to expand parent nodes")
	}

	// Test FindAndSelect with non-existent node
	if tv.FindAndSelect("NonExistent") {
		t.Error("FindAndSelect should return false for non-existent node")
	}
}

// TestTreeViewStyle tests style functionality
func TestTreeViewStyle(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)
	root := createSampleTree()
	tv.SetRoot(root)

	// Test DefaultTreeViewStyle
	style := DefaultTreeViewStyle()
	if style.ExpandedIcon != '-' || style.CollapsedIcon != '+' {
		t.Error("DefaultTreeViewStyle returned incorrect icons")
	}

	// Test SetStyle
	newStyle := DefaultTreeViewStyle()
	newStyle.NodeStyle = tcell.StyleDefault.Foreground(tcell.ColorRed)
	tv.SetStyle(newStyle)
	if tv.Style != newStyle.NodeStyle {
		t.Error("SetStyle failed to update widget style")
	}

	// Test SetNodeStyle
	customStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue)
	tv.SetNodeStyle(root, customStyle)
	if root.Style != customStyle {
		t.Error("SetNodeStyle failed to update node style")
	}
}

// TestTreeViewEdgeCases tests edge cases and error conditions
func TestTreeViewEdgeCases(t *testing.T) {
	screen := createTestScreen(t)
	defer screen.Fini()

	tv := NewTreeView(screen, 0, 0, 40, 10)

	// Test operations on empty tree
	if tv.SelectNext() {
		t.Error("SelectNext should return false on empty tree")
	}
	if tv.SelectPrevious() {
		t.Error("SelectPrevious should return false on empty tree")
	}
	if tv.ToggleSelected() {
		t.Error("ToggleSelected should return false on empty tree")
	}
	if tv.FindNode("anything") != nil {
		t.Error("FindNode should return nil on empty tree")
	}

	// Test AddNode with nil parent (should create root)
	root := tv.AddNode(nil, "Root")
	if root == nil {
		t.Error("AddNode with nil parent should create root node")
	}
	if tv.Root != root {
		t.Error("AddNode with nil parent failed to set root")
	}

	// Test adding second root (should fail)
	secondRoot := tv.AddNode(nil, "Second Root")
	if secondRoot != nil {
		t.Error("Adding second root node should fail")
	}

	// Test RemoveNode edge cases
	tv.RemoveNode(nil) // Should not panic
	tv.RemoveNode(root)
	if tv.Root != nil {
		t.Error("RemoveNode failed to clear root")
	}
	if tv.Selected != nil {
		t.Error("RemoveNode failed to clear selection")
	}
}
