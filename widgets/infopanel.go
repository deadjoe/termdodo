package widgets

import (
	"strings"

	"github.com/deadjoe/termdodo/theme"
	"github.com/gdamore/tcell/v2"
)

// InfoField represents a field in the info panel
type InfoField struct {
	Label      string
	Value      string
	LabelStyle tcell.Style
	ValueStyle tcell.Style
}

// InfoPanel represents an info panel widget
type InfoPanel struct {
	X, Y          int
	Width, Height int
	Screen        tcell.Screen
	Style         tcell.Style
	TitleStyle    tcell.Style

	Title        string
	Fields       []InfoField
	ShowBorder   bool
	ScrollOffset int
	LabelWidth   int
}

// NewInfoPanel creates a new info panel widget
func NewInfoPanel(screen tcell.Screen, x, y, width, height int) *InfoPanel {
	return &InfoPanel{
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		Screen:     screen,
		Style:      theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		TitleStyle: theme.GetStyle(theme.Current.Title, theme.Current.MainBg),
		ShowBorder: true,
		LabelWidth: 20,
	}
}

// SetTitle sets the panel title
func (p *InfoPanel) SetTitle(title string) {
	p.Title = title
}

// SetFields sets the info fields
func (p *InfoPanel) SetFields(fields []InfoField) {
	p.Fields = fields
}

// AddField adds a field to the panel
func (p *InfoPanel) AddField(label, value string) {
	p.Fields = append(p.Fields, InfoField{
		Label: label,
		Value: value,
	})
}

// ClearFields clears all fields
func (p *InfoPanel) ClearFields() {
	p.Fields = nil
}

// Draw draws the info panel
func (p *InfoPanel) Draw() {
	if p.ShowBorder {
		p.drawBorder()
	}

	// Draw title if present
	startY := p.Y
	if p.ShowBorder {
		startY++
	}
	if p.Title != "" {
		p.drawTitle(startY)
		startY++
	}

	// Draw fields
	p.drawFields(startY)
}

// drawBorder draws the panel border
func (p *InfoPanel) drawBorder() {
	style := p.Style

	// Draw corners
	p.Screen.SetContent(p.X, p.Y, '┌', nil, style)
	p.Screen.SetContent(p.X+p.Width-1, p.Y, '┐', nil, style)
	p.Screen.SetContent(p.X, p.Y+p.Height-1, '└', nil, style)
	p.Screen.SetContent(p.X+p.Width-1, p.Y+p.Height-1, '┘', nil, style)

	// Draw horizontal lines
	for x := p.X + 1; x < p.X+p.Width-1; x++ {
		p.Screen.SetContent(x, p.Y, '─', nil, style)
		p.Screen.SetContent(x, p.Y+p.Height-1, '─', nil, style)
	}

	// Draw vertical lines
	for y := p.Y + 1; y < p.Y+p.Height-1; y++ {
		p.Screen.SetContent(p.X, y, '│', nil, style)
		p.Screen.SetContent(p.X+p.Width-1, y, '│', nil, style)
	}
}

// drawTitle draws the panel title
func (p *InfoPanel) drawTitle(y int) {
	if p.Title == "" {
		return
	}

	x := p.X
	if p.ShowBorder {
		x++
	}

	// Center the title
	titleWidth := len(p.Title)
	if titleWidth > p.Width-2 {
		titleWidth = p.Width - 2
	}
	startX := x + (p.Width-2-titleWidth)/2

	// Draw title
	for i, r := range p.Title {
		if i >= titleWidth {
			break
		}
		p.Screen.SetContent(startX+i, y, r, nil, p.TitleStyle)
	}
}

// drawFields draws the info fields
func (p *InfoPanel) drawFields(startY int) {
	visibleHeight := p.Height
	if p.ShowBorder {
		visibleHeight -= 2
	}
	if p.Title != "" {
		visibleHeight--
	}

	x := p.X
	if p.ShowBorder {
		x++
	}

	// Draw fields
	for i := p.ScrollOffset; i < len(p.Fields) && i-p.ScrollOffset < visibleHeight; i++ {
		field := p.Fields[i]
		y := startY + i - p.ScrollOffset

		// Draw label
		labelStyle := field.LabelStyle
		if labelStyle == (tcell.Style{}) {
			labelStyle = p.Style
		}
		label := field.Label + ":"
		if len(label) > p.LabelWidth {
			label = label[:p.LabelWidth-3] + "..."
		} else {
			label = label + strings.Repeat(" ", p.LabelWidth-len(label))
		}
		for i, r := range label {
			p.Screen.SetContent(x+i, y, r, nil, labelStyle)
		}

		// Draw value
		valueStyle := field.ValueStyle
		if valueStyle == (tcell.Style{}) {
			valueStyle = p.Style
		}
		value := field.Value
		maxValueWidth := p.Width - p.LabelWidth - 3
		if len(value) > maxValueWidth {
			value = value[:maxValueWidth-3] + "..."
		}
		for i, r := range value {
			p.Screen.SetContent(x+p.LabelWidth+1+i, y, r, nil, valueStyle)
		}
	}
}

// SetLabelWidth sets the width of the label column
func (p *InfoPanel) SetLabelWidth(width int) {
	p.LabelWidth = width
}

// SetShowBorder sets whether to show the border
func (p *InfoPanel) SetShowBorder(show bool) {
	p.ShowBorder = show
}

// HandleEvent handles keyboard events for scrolling
func (p *InfoPanel) HandleEvent(ev *tcell.EventKey) bool {
	visibleHeight := p.Height
	if p.ShowBorder {
		visibleHeight -= 2
	}
	if p.Title != "" {
		visibleHeight--
	}

	switch ev.Key() {
	case tcell.KeyUp:
		if p.ScrollOffset > 0 {
			p.ScrollOffset--
			return true
		}
	case tcell.KeyDown:
		if p.ScrollOffset < len(p.Fields)-visibleHeight {
			p.ScrollOffset++
			return true
		}
	case tcell.KeyHome:
		if p.ScrollOffset > 0 {
			p.ScrollOffset = 0
			return true
		}
	case tcell.KeyEnd:
		maxOffset := len(p.Fields) - visibleHeight
		if maxOffset > 0 && p.ScrollOffset < maxOffset {
			p.ScrollOffset = maxOffset
			return true
		}
	}
	return false
}

// GetHeight returns the total height of the panel
func (p *InfoPanel) GetHeight() int {
	return p.Height
}

// GetWidth returns the total width of the panel
func (p *InfoPanel) GetWidth() int {
	return p.Width
}

// SetStyle sets the default style for the panel
func (p *InfoPanel) SetStyle(style tcell.Style) {
	p.Style = style
}

// SetTitleStyle sets the style for the title
func (p *InfoPanel) SetTitleStyle(style tcell.Style) {
	p.TitleStyle = style
}
