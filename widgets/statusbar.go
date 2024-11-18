package widgets

import (
	"github.com/gdamore/tcell/v2"
	"termdodo/theme"
	"strings"
)

// StatusItem represents a single item in the status bar
type StatusItem struct {
	Text      string
	Style     tcell.Style
	MinWidth  int
	MaxWidth  int
	Alignment Alignment
}

// StatusBar represents a status bar widget
type StatusBar struct {
	X, Y    int
	Width   int
	Screen  tcell.Screen
	Style   tcell.Style
	Items   []StatusItem
	Padding int
}

// NewStatusBar creates a new status bar widget
func NewStatusBar(screen tcell.Screen, x, y, width int) *StatusBar {
	return &StatusBar{
		X:       x,
		Y:       y,
		Width:   width,
		Screen:  screen,
		Style:   theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		Padding: 1,
	}
}

// SetItems sets the status bar items
func (s *StatusBar) SetItems(items []StatusItem) {
	s.Items = items
}

// AddItem adds an item to the status bar
func (s *StatusBar) AddItem(item StatusItem) {
	s.Items = append(s.Items, item)
}

// ClearItems clears all items from the status bar
func (s *StatusBar) ClearItems() {
	s.Items = nil
}

// Draw draws the status bar
func (s *StatusBar) Draw() {
	if len(s.Items) == 0 {
		return
	}

	// Calculate total minimum width and flexible width
	totalMinWidth := 0
	totalFlexWidth := 0
	for _, item := range s.Items {
		minWidth := item.MinWidth
		if minWidth == 0 {
			minWidth = len(item.Text) + s.Padding*2
		}
		totalMinWidth += minWidth
		if item.MaxWidth > minWidth {
			totalFlexWidth += item.MaxWidth - minWidth
		}
	}

	// Calculate available width for flexible items
	availableWidth := s.Width
	extraWidth := availableWidth - totalMinWidth

	// Distribute extra width among flexible items
	if extraWidth > 0 && totalFlexWidth > 0 {
		for i := range s.Items {
			if s.Items[i].MaxWidth > s.Items[i].MinWidth {
				flexRatio := float64(s.Items[i].MaxWidth-s.Items[i].MinWidth) / float64(totalFlexWidth)
				extra := int(float64(extraWidth) * flexRatio)
				if s.Items[i].MinWidth+extra > s.Items[i].MaxWidth {
					extra = s.Items[i].MaxWidth - s.Items[i].MinWidth
				}
				s.Items[i].MinWidth += extra
			}
		}
	}

	// Draw items
	x := s.X
	for _, item := range s.Items {
		text := item.Text
		style := item.Style
		if style == (tcell.Style{}) {
			style = s.Style
		}

		// Add padding
		text = strings.Repeat(" ", s.Padding) + text + strings.Repeat(" ", s.Padding)

		// Truncate or pad text to fit width
		width := item.MinWidth
		if len(text) > width {
			text = text[:width]
		} else if len(text) < width {
			switch item.Alignment {
			case AlignLeft:
				text = text + strings.Repeat(" ", width-len(text))
			case AlignRight:
				text = strings.Repeat(" ", width-len(text)) + text
			case AlignCenter:
				leftPad := (width - len(text)) / 2
				rightPad := width - len(text) - leftPad
				text = strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
			}
		}

		// Draw text
		for i, r := range text {
			s.Screen.SetContent(x+i, s.Y, r, nil, style)
		}
		x += width
	}
}

// SetStyle sets the default style for the status bar
func (s *StatusBar) SetStyle(style tcell.Style) {
	s.Style = style
}

// SetPadding sets the padding between items
func (s *StatusBar) SetPadding(padding int) {
	s.Padding = padding
}

// GetWidth returns the total width of all items
func (s *StatusBar) GetWidth() int {
	width := 0
	for _, item := range s.Items {
		width += item.MinWidth
	}
	return width
}

// GetHeight returns the height of the status bar (always 1)
func (s *StatusBar) GetHeight() int {
	return 1
}
