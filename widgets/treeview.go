package widgets

import (
	theme "github.com/deadjoe/termdodo/theme"
	tcell "github.com/gdamore/tcell/v2"
)

// TreeNode represents a node in the tree
type TreeNode struct {
	Text     string
	Children []*TreeNode
	Parent   *TreeNode
	Expanded bool
	Data     interface{} // Optional data associated with the node
	Style    tcell.Style
}

// TreeView represents a tree view widget
type TreeView struct {
	X, Y          int
	Width, Height int
	Screen        tcell.Screen
	Style         tcell.Style
	SelectedStyle tcell.Style

	Root         *TreeNode
	Selected     *TreeNode
	ScrollOffset int
	VisibleNodes int

	ShowLines bool
	Indent    int
}

// TreeViewStyle represents the style configuration for the tree view
type TreeViewStyle struct {
	NodeStyle      tcell.Style // Style for normal nodes
	SelectedStyle  tcell.Style // Style for selected node
	LineStyle      tcell.Style // Style for tree lines
	ExpandedIcon   rune        // Icon for expanded nodes
	CollapsedIcon  rune        // Icon for collapsed nodes
}

// DefaultTreeViewStyle returns the default style configuration
func DefaultTreeViewStyle() TreeViewStyle {
	return TreeViewStyle{
		NodeStyle:     theme.GetStyle(theme.ColorToHex(theme.Current.MainFg), theme.ColorToHex(theme.Current.MainBg)),
		SelectedStyle: theme.GetStyle(theme.ColorToHex(theme.Current.Selected), theme.ColorToHex(theme.Current.HighlightBg)),
		LineStyle:     theme.GetStyle(theme.ColorToHex(theme.Current.MainFg), theme.ColorToHex(theme.Current.MainBg)),
		ExpandedIcon:  '-',
		CollapsedIcon: '+',
	}
}

// NewTreeView creates a new tree view widget
func NewTreeView(screen tcell.Screen, x, y, width, height int) *TreeView {
	return &TreeView{
		X:             x,
		Y:             y,
		Width:         width,
		Height:        height,
		Screen:        screen,
		Style:         theme.GetStyle(theme.ColorToHex(theme.Current.MainFg), theme.ColorToHex(theme.Current.MainBg)),
		SelectedStyle: theme.GetStyle(theme.ColorToHex(theme.Current.Selected), theme.ColorToHex(theme.Current.HighlightBg)),
		ShowLines:     true,
		Indent:        2,
	}
}

// SetRoot sets the root node of the tree
func (t *TreeView) SetRoot(root *TreeNode) {
	t.Root = root
	t.Selected = nil
	t.ScrollOffset = 0
}

// Draw draws the tree view
func (t *TreeView) Draw() {
	if t.Root == nil {
		return
	}

	t.VisibleNodes = 0
	t.drawNode(t.Root, t.X, t.Y-t.ScrollOffset, false)
}

// getNodeStyle returns the appropriate style for a node
func (t *TreeView) getNodeStyle(node *TreeNode) tcell.Style {
	style := t.Style
	if node.Style != (tcell.Style{}) {
		style = node.Style
	}
	if node == t.Selected {
		style = style.Reverse(true)
	}
	return style
}

// drawNode recursively draws a node and its children
func (t *TreeView) drawNode(node *TreeNode, x, y int, isLast bool) int {
	if y >= t.Y+t.Height {
		return y
	}

	if y >= t.Y {
		t.VisibleNodes++
		// Draw node content
		style := t.getNodeStyle(node)

		// Draw tree lines
		if t.ShowLines && node != t.Root {
			for i := t.X; i < x-t.Indent; i += t.Indent {
				t.Screen.SetContent(i, y, '│', nil, t.Style)
			}
			if isLast {
				t.Screen.SetContent(x-t.Indent, y, '└', nil, t.Style)
			} else {
				t.Screen.SetContent(x-t.Indent, y, '├', nil, t.Style)
			}
			for i := x - t.Indent + 1; i < x; i++ {
				t.Screen.SetContent(i, y, '─', nil, t.Style)
			}
		}

		// Draw expand/collapse indicator
		if len(node.Children) > 0 {
			if node.Expanded {
				t.Screen.SetContent(x, y, '-', nil, style)
			} else {
				t.Screen.SetContent(x, y, '+', nil, style)
			}
			x += 2
		}

		// Draw node text
		for i, r := range node.Text {
			if x+i >= t.X+t.Width {
				break
			}
			t.Screen.SetContent(x+i, y, r, nil, style)
		}
	}

	y++

	if node.Expanded {
		childX := x + t.Indent
		for i, child := range node.Children {
			isLastChild := i == len(node.Children)-1
			y = t.drawNode(child, childX, y, isLastChild)
		}
	}

	return y
}

// HandleKeyEvent handles keyboard events for the tree view
func (t *TreeView) HandleKeyEvent(event *tcell.EventKey) bool {
	if t.Root == nil {
		return false
	}

	if t.Selected == nil {
		t.Selected = t.Root
	}

	switch event.Key() {
	case tcell.KeyUp:
		return t.SelectPrevious()
	case tcell.KeyDown:
		return t.SelectNext()
	case tcell.KeyLeft:
		if t.Selected.Expanded {
			t.Selected.Expanded = false
			return true
		} else if t.Selected.Parent != nil {
			t.Selected = t.Selected.Parent
			return true
		}
	case tcell.KeyRight:
		if !t.Selected.Expanded && len(t.Selected.Children) > 0 {
			t.Selected.Expanded = true
			return true
		} else if t.Selected.Expanded && len(t.Selected.Children) > 0 {
			t.Selected = t.Selected.Children[0]
			return true
		}
	}
	return false
}

// isLeafNode returns true if the node is a leaf node
func (n *TreeNode) isLeafNode() bool {
	return len(n.Children) == 0
}

// SelectNext selects the next visible node in the tree
func (t *TreeView) SelectNext() bool {
	if t.Root == nil || t.Selected == nil {
		if t.Root != nil {
			t.Selected = t.Root
			return true
		}
		return false
	}

	// If we're at a leaf node, we can't go further down
	if t.Selected.isLeafNode() {
		return false
	}

	// If current node is expanded and has children, select first child
	if t.Selected.Expanded && len(t.Selected.Children) > 0 {
		t.Selected = t.Selected.Children[0]
		return true
	}

	// Try to find next sibling or ancestor's sibling
	current := t.Selected
	for current != nil {
		parent := current.Parent
		if parent == nil {
			return false // At root with no children or not expanded
		}

		// Find current node's index among siblings
		for i, sibling := range parent.Children {
			if sibling == current {
				// If there's a next sibling, select it
				if i < len(parent.Children)-1 {
					t.Selected = parent.Children[i+1]
					return true
				}
				// No more siblings, move up to parent and continue
				current = parent
				break
			}
		}
	}

	// No more nodes to select, keep current selection
	return false
}

// SelectPrevious selects the previous visible node
func (t *TreeView) SelectPrevious() bool {
	if t.Root == nil || t.Selected == nil {
		if t.Root != nil {
			t.Selected = t.Root
			return true
		}
		return false
	}

	// If at root, cannot go up further
	if t.Selected == t.Root {
		return false
	}

	parent := t.Selected.Parent
	if parent == nil {
		return false
	}

	// Find current node's position among siblings
	for i, sibling := range parent.Children {
		if sibling == t.Selected {
			if i > 0 {
				// Select the deepest visible node in the previous sibling's subtree
				prev := parent.Children[i-1]
				for prev.Expanded && len(prev.Children) > 0 {
					prev = prev.Children[len(prev.Children)-1]
				}
				t.Selected = prev
			} else {
				// If first child, select parent
				t.Selected = parent
			}
			return true
		}
	}

	return false
}

// GetSelected returns the currently selected node
func (t *TreeView) GetSelected() *TreeNode {
	return t.Selected
}

// SetShowLines sets whether to show tree lines
func (t *TreeView) SetShowLines(show bool) {
	t.ShowLines = show
}

// SetIndent sets the indentation level
func (t *TreeView) SetIndent(indent int) {
	if indent < 0 {
		indent = 0
	}
	t.Indent = indent
}

// ExpandAll expands all nodes in the tree
func (t *TreeView) ExpandAll() {
	t.expandNode(t.Root)
}

// CollapseAll collapses all nodes in the tree
func (t *TreeView) CollapseAll() {
	t.collapseNode(t.Root)
}

// expandNode recursively expands a node and its children
func (t *TreeView) expandNode(node *TreeNode) {
	if node == nil {
		return
	}
	node.Expanded = true
	for _, child := range node.Children {
		t.expandNode(child)
	}
}

// collapseNode recursively collapses a node and its children
func (t *TreeView) collapseNode(node *TreeNode) {
	if node == nil {
		return
	}
	node.Expanded = false
	for _, child := range node.Children {
		t.collapseNode(child)
	}
}

// AddNode adds a child node to the specified parent
// If parent is nil, the node will be set as the root node
func (t *TreeView) AddNode(parent *TreeNode, text string) *TreeNode {
	node := &TreeNode{
		Text:   text,
		Style:  t.Style,
	}

	if parent == nil {
		// Set as root node
		if t.Root != nil {
			// If root exists, return nil to prevent multiple roots
			return nil
		}
		t.Root = node
	} else {
		node.Parent = parent
		parent.Children = append(parent.Children, node)
	}

	// Select the new node if no node is currently selected
	if t.Selected == nil {
		t.Selected = node
	}

	return node
}

// RemoveNode removes a node and its children from the tree
func (t *TreeView) RemoveNode(node *TreeNode) {
	if node == nil {
		return
	}

	// If removing the root node
	if node == t.Root {
		t.Root = nil
		t.Selected = nil
		return
	}

	// Find and remove the node from its parent's children
	if node.Parent != nil {
		children := node.Parent.Children
		for i, child := range children {
			if child == node {
				// Remove the node from its parent's children
				node.Parent.Children = append(children[:i], children[i+1:]...)

				// Clear the parent reference
				node.Parent = nil

				// If the removed node was selected, select the parent
				if node == t.Selected {
					t.Selected = node.Parent
				}

				break
			}
		}
	}

	// Clear all parent references in the removed subtree
	var clearParents func(*TreeNode)
	clearParents = func(n *TreeNode) {
		if n == nil {
			return
		}
		n.Parent = nil
		for _, child := range n.Children {
			clearParents(child)
		}
	}
	clearParents(node)
}

// ExpandSelected expands the currently selected node
func (t *TreeView) ExpandSelected() bool {
	if t.Selected != nil {
		t.Selected.Expanded = true
		return true
	}
	return false
}

// CollapseSelected collapses the currently selected node
func (t *TreeView) CollapseSelected() bool {
	if t.Selected != nil {
		t.Selected.Expanded = false
		return true
	}
	return false
}

// ToggleSelected toggles the expanded state of the currently selected node
func (t *TreeView) ToggleSelected() bool {
	if t.Selected != nil {
		t.Selected.Expanded = !t.Selected.Expanded
		return true
	}
	return false
}

// EnsureVisible ensures the selected node is visible in the view
func (t *TreeView) EnsureVisible() {
	if t.Selected == nil {
		return
	}

	// First ensure all parent nodes are expanded
	current := t.Selected.Parent
	for current != nil {
		current.Expanded = true
		current = current.Parent
	}

	// Calculate node position
	y := 0
	var calcPosition func(*TreeNode) int
	calcPosition = func(node *TreeNode) int {
		if node == t.Selected {
			return y
		}
		if node == nil {
			return -1
		}
		y++
		if node.Expanded {
			for _, child := range node.Children {
				if result := calcPosition(child); result >= 0 {
					return result
				}
			}
		}
		return -1
	}

	pos := calcPosition(t.Root)
	if pos < 0 {
		return
	}

	// Adjust scroll offset if needed
	if pos < t.ScrollOffset {
		t.ScrollOffset = pos
	} else if pos >= t.ScrollOffset+t.Height {
		t.ScrollOffset = pos - t.Height + 1
	}

	if t.ScrollOffset < 0 {
		t.ScrollOffset = 0
	}
}

// ScrollTo scrolls the view to the specified offset
func (t *TreeView) ScrollTo(offset int) {
	if offset < 0 {
		offset = 0
	}
	t.ScrollOffset = offset
}

// ScrollBy scrolls the view by the specified amount
func (t *TreeView) ScrollBy(delta int) {
	t.ScrollTo(t.ScrollOffset + delta)
}

// SetStyle sets the style configuration for the tree view
func (t *TreeView) SetStyle(style TreeViewStyle) {
	t.Style = style.NodeStyle
	t.SelectedStyle = style.SelectedStyle
	// Update all existing nodes with the new style
	var updateStyles func(*TreeNode)
	updateStyles = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Style == (tcell.Style{}) {
			node.Style = style.NodeStyle
		}
		for _, child := range node.Children {
			updateStyles(child)
		}
	}
	updateStyles(t.Root)
}

// SetNodeStyle sets the style for a specific node
func (t *TreeView) SetNodeStyle(node *TreeNode, style tcell.Style) {
	if node != nil {
		node.Style = style
	}
}

// FindNode searches for a node with the given text
func (t *TreeView) FindNode(text string) *TreeNode {
	if t.Root == nil {
		return nil
	}

	var search func(*TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Text == text {
			return node
		}

		for _, child := range node.Children {
			if found := search(child); found != nil {
				return found
			}
		}

		return nil
	}

	return search(t.Root)
}

// FindAndSelect searches for a node with the given text and selects it if found
func (t *TreeView) FindAndSelect(text string) bool {
	node := t.FindNode(text)
	if node != nil {
		// Expand all parent nodes to make the found node visible
		current := node.Parent
		for current != nil {
			current.Expanded = true
			current = current.Parent
		}
		t.Selected = node
		t.EnsureVisible()
		return true
	}
	return false
}
