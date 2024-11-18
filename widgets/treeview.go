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
	t.drawNode(t.Root, t.X, t.Y, false)
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

// SelectNext selects the next visible node in the tree
func (t *TreeView) SelectNext() bool {
	if t.Root == nil {
		return false
	}

	if t.Selected == nil {
		t.Selected = t.Root
		return true
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
			return false
		}

		siblings := parent.Children
		for i, sibling := range siblings {
			if sibling == current && i < len(siblings)-1 {
				t.Selected = siblings[i+1]
				return true
			}
		}
		current = parent
	}

	return false
}

// SelectPrevious selects the previous visible node
func (t *TreeView) SelectPrevious() bool {
	if t.Root == nil {
		return false
	}

	if t.Selected == nil {
		t.Selected = t.Root
		return true
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
				// Select the previous sibling's deepest visible node
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
func (t *TreeView) AddNode(parent *TreeNode, text string) *TreeNode {
	if parent == nil {
		return nil
	}
	node := &TreeNode{
		Text:   text,
		Parent: parent,
		Style:  t.Style,
	}
	parent.Children = append(parent.Children, node)
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
