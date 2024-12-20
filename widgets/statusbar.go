package widgets

import (
	"github.com/deadjoe/termdodo/theme"
	"github.com/gdamore/tcell/v2"
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
	X, Y          int
	Width, Height int
	Screen        tcell.Screen
	Style         tcell.Style

	Items     []StatusItem
	Separator string
	Padding   int
}

// NewStatusBar creates a new status bar widget
func NewStatusBar(screen tcell.Screen, x, y, width int) *StatusBar {
	return &StatusBar{
		X:         x,
		Y:         y,
		Width:     width,
		Height:    1,
		Screen:    screen,
		Style:     theme.Current.GetStyle(),
		Items:     nil,
		Separator: " | ",
		Padding:   1,
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

// UpdateItem updates an item at the specified index
func (s *StatusBar) UpdateItem(index int, text string, style ...tcell.Style) {
	if index >= 0 && index < len(s.Items) {
		s.Items[index].Text = text
		if len(style) > 0 {
			s.Items[index].Style = style[0]
		}
	}
}

// calculateItemWidths calculates the minimum and flexible widths for all items
func (s *StatusBar) calculateItemWidths() (int, int) {
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
	return totalMinWidth, totalFlexWidth
}

// distributeExtraWidth distributes extra width among flexible items
func (s *StatusBar) distributeExtraWidth(extraWidth, totalFlexWidth int) {
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
}

// drawItem draws a single status bar item
func (s *StatusBar) drawItem(x int, item StatusItem) int {
	text := item.Text
	width := item.MinWidth
	if width == 0 {
		width = len(text) + s.Padding*2
	}

	// Add padding
	if s.Padding > 0 {
		text = strings.Repeat(" ", s.Padding) + text + strings.Repeat(" ", s.Padding)
	}

	// Truncate or pad text to fit width
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
	style := item.Style
	if style == tcell.StyleDefault {
		style = s.Style
	}
	for i, r := range text {
		s.Screen.SetContent(x+i, s.Y, r, nil, style)
	}

	return width
}

// Draw draws the status bar
func (s *StatusBar) Draw() {
	if len(s.Items) == 0 {
		return
	}

	// Calculate widths
	totalMinWidth, totalFlexWidth := s.calculateItemWidths()

	// Calculate and distribute extra width
	availableWidth := s.Width
	extraWidth := availableWidth - totalMinWidth - (len(s.Items)-1)*len(s.Separator)
	s.distributeExtraWidth(extraWidth, totalFlexWidth)

	// Draw items
	x := s.X
	for i, item := range s.Items {
		width := s.drawItem(x, item)
		x += width
		if i < len(s.Items)-1 {
			for j, r := range s.Separator {
				s.Screen.SetContent(x+j, s.Y, r, nil, s.Style)
			}
			x += len(s.Separator)
		}
	}
}

// SetStyle sets the default style for the status bar
func (s *StatusBar) SetStyle(style tcell.Style) {
	s.Style = style
}

// SetSeparator sets the separator between status items
func (s *StatusBar) SetSeparator(sep string) {
	s.Separator = sep
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
	return width + (len(s.Items)-1)*len(s.Separator)
}

// GetHeight returns the height of the status bar (always 1)
func (s *StatusBar) GetHeight() int {
	return 1
}
