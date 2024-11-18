package widgets

import (
	"github.com/gdamore/tcell/v2"
	"termdodo/theme"
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

	Root        *TreeNode
	Selected    *TreeNode
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
		Style:         theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		SelectedStyle: theme.GetStyle(theme.Current.Selected, theme.Current.HighlightBg),
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

	// Reset visible nodes count
	t.VisibleNodes = 0

	// Draw nodes starting from root
	t.drawNode(t.Root, t.X, t.Y, 0)
}

// drawNode recursively draws a node and its children
func (t *TreeView) drawNode(node *TreeNode, x, y, level int) {
	if t.VisibleNodes-t.ScrollOffset >= t.Height {
		return
	}

	if t.VisibleNodes >= t.ScrollOffset {
		// Calculate actual y position
		actualY := y + t.VisibleNodes - t.ScrollOffset

		// Draw node
		style := t.Style
		if node == t.Selected {
			style = t.SelectedStyle
		} else if node.Style != (tcell.Style{}) {
			style = node.Style
		}

		// Draw tree lines if enabled
		if t.ShowLines && level > 0 {
			lineX := x
			for i := 0; i < level*t.Indent; i++ {
				char := ' '
				if i%t.Indent == 0 && i < (level-1)*t.Indent {
					char = '│'
				}
				t.Screen.SetContent(lineX+i, actualY, char, nil, t.Style)
			}
			
			// Draw connection to parent
			if node.Parent != nil {
				isLast := node == node.Parent.Children[len(node.Parent.Children)-1]
				if isLast {
					t.Screen.SetContent(x+(level-1)*t.Indent, actualY, '└', nil, t.Style)
				} else {
					t.Screen.SetContent(x+(level-1)*t.Indent, actualY, '├', nil, t.Style)
				}
			}
		}

		// Draw expand/collapse indicator
		nodeX := x + level*t.Indent
		if len(node.Children) > 0 {
			if node.Expanded {
				t.Screen.SetContent(nodeX, actualY, '▼', nil, style)
			} else {
				t.Screen.SetContent(nodeX, actualY, '▶', nil, style)
			}
			nodeX += 2
		} else {
			nodeX += 2
		}

		// Draw node text
		for i, r := range []rune(node.Text) {
			if nodeX+i >= t.X+t.Width {
				break
			}
			t.Screen.SetContent(nodeX+i, actualY, r, nil, style)
		}
	}
	t.VisibleNodes++

	// Draw children if expanded
	if node.Expanded {
		for _, child := range node.Children {
			t.drawNode(child, x, y, level+1)
		}
	}
}

// HandleEvent handles keyboard events
func (t *TreeView) HandleEvent(ev *tcell.EventKey) bool {
	if t.Root == nil {
		return false
	}

	switch ev.Key() {
	case tcell.KeyUp:
		return t.selectPrevious()
	case tcell.KeyDown:
		return t.selectNext()
	case tcell.KeyRight:
		if t.Selected != nil {
			t.Selected.Expanded = true
			return true
		}
	case tcell.KeyLeft:
		if t.Selected != nil {
			if t.Selected.Expanded && len(t.Selected.Children) > 0 {
				t.Selected.Expanded = false
			} else if t.Selected.Parent != nil {
				t.Selected = t.Selected.Parent
			}
			return true
		}
	}
	return false
}

// selectNext selects the next visible node
func (t *TreeView) selectNext() bool {
	if t.Selected == nil {
		t.Selected = t.Root
		return true
	}

	// If current node is expanded and has children, select first child
	if t.Selected.Expanded && len(t.Selected.Children) > 0 {
		t.Selected = t.Selected.Children[0]
		return true
	}

	// Otherwise, find next sibling or parent's next sibling
	current := t.Selected
	for current != nil {
		if current.Parent == nil {
			return false
		}

		siblings := current.Parent.Children
		for i, sibling := range siblings {
			if sibling == current && i < len(siblings)-1 {
				t.Selected = siblings[i+1]
				return true
			}
		}
		current = current.Parent
	}

	return false
}

// selectPrevious selects the previous visible node
func (t *TreeView) selectPrevious() bool {
	if t.Selected == nil || t.Selected == t.Root {
		return false
	}

	siblings := t.Selected.Parent.Children
	for i, sibling := range siblings {
		if sibling == t.Selected {
			if i > 0 {
				// Select the last visible child of the previous sibling
				t.Selected = t.getLastVisibleNode(siblings[i-1])
			} else {
				// Select parent
				t.Selected = t.Selected.Parent
			}
			return true
		}
	}

	return false
}

// getLastVisibleNode returns the last visible node in a subtree
func (t *TreeView) getLastVisibleNode(node *TreeNode) *TreeNode {
	if !node.Expanded || len(node.Children) == 0 {
		return node
	}
	return t.getLastVisibleNode(node.Children[len(node.Children)-1])
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
	node := &TreeNode{
		Text:     text,
		Parent:   parent,
		Expanded: false,
	}
	if parent == nil {
		t.Root = node
	} else {
		parent.Children = append(parent.Children, node)
	}
	return node
}

// RemoveNode removes a node and its children from the tree
func (t *TreeView) RemoveNode(node *TreeNode) {
	if node == nil || node.Parent == nil {
		return
	}

	siblings := node.Parent.Children
	for i, sibling := range siblings {
		if sibling == node {
			node.Parent.Children = append(siblings[:i], siblings[i+1:]...)
			break
		}
	}
}
