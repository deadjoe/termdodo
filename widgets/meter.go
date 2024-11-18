package widgets

import (
	"fmt"

	"github.com/deadjoe/termdodo/theme"
	tcell "github.com/gdamore/tcell/v2"
)

// Meter represents a percentage meter widget
type Meter struct {
	X, Y    int
	Width   int
	Style   tcell.Style
	Screen  tcell.Screen
	Value   float64
	ShowPct bool
	Label   string

	// New fields for block style and gradient
	BlockStyle   bool
	BlockSpacing int
	StartColor   tcell.Color
	EndColor     tcell.Color
	UseGradient  bool
}

// NewMeter creates a new meter widget
func NewMeter(screen tcell.Screen, x, y, width int) *Meter {
	return &Meter{
		X:            x,
		Y:            y,
		Width:        width,
		Style:        theme.Current.GetStyle(),
		Screen:       screen,
		Value:        0,
		ShowPct:      true,
		Label:        "",
		BlockStyle:   false,
		BlockSpacing: 0,
		StartColor:   theme.Current.Foreground,
		EndColor:     theme.Current.Foreground,
		UseGradient:  false,
	}
}

// SetBlockStyle sets whether to use block style display
func (m *Meter) SetBlockStyle(enabled bool) {
	m.BlockStyle = enabled
}

// SetBlockSpacing sets the spacing between blocks
func (m *Meter) SetBlockSpacing(spacing int) {
	m.BlockSpacing = spacing
}

// SetGradient sets the start and end colors for gradient
func (m *Meter) SetGradient(start, end tcell.Color) {
	m.StartColor = start
	m.EndColor = end
	m.UseGradient = true
}

// Draw draws the meter on the screen
func (m *Meter) Draw() {
	// Calculate filled width
	filledWidth := int(float64(m.Width) * m.Value)
	if filledWidth > m.Width {
		filledWidth = m.Width
	}

	// Draw the meter
	if m.BlockStyle {
		// Block style
		blockWidth := 1
		if m.BlockSpacing > 0 {
			blockWidth += m.BlockSpacing
		}
		numBlocks := m.Width / blockWidth
		filledBlocks := int(float64(numBlocks) * m.Value)

		for i := 0; i < numBlocks; i++ {
			x := m.X + i*blockWidth
			var style tcell.Style
			if i < filledBlocks {
				if m.UseGradient {
					position := float64(i) / float64(numBlocks-1)
					style = tcell.StyleDefault.
						Background(m.StartColor).
						Foreground(interpolateColor(m.StartColor, m.EndColor, position))
				} else {
					style = m.Style
				}
			} else {
				style = theme.Current.GetStyle()
			}
			m.Screen.SetContent(x, m.Y, '█', nil, style)
		}
	} else {
		// Regular style
		for i := 0; i < m.Width; i++ {
			var style tcell.Style
			if i < filledWidth {
				if m.UseGradient {
					position := float64(i) / float64(m.Width-1)
					style = tcell.StyleDefault.
						Background(m.StartColor).
						Foreground(interpolateColor(m.StartColor, m.EndColor, position))
				} else {
					style = m.Style
				}
			} else {
				style = theme.Current.GetStyle()
			}
			m.Screen.SetContent(m.X+i, m.Y, '█', nil, style)
		}
	}

	// Draw percentage if enabled
	if m.ShowPct {
		text := fmt.Sprintf("%3.0f%%", m.Value*100)
		textStyle := theme.Current.GetStyle()
		m.drawTextStyled(m.X+m.Width+1, m.Y, text, textStyle)
	}
}

// SetValue sets the current value of the meter (0-1)
func (m *Meter) SetValue(value float64) {
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	m.Value = value
}

// SetStyle sets the style for the meter
func (m *Meter) SetStyle(style tcell.Style) {
	m.Style = style
}

// SetShowPercentage sets whether to show the percentage value
func (m *Meter) SetShowPercentage(show bool) {
	m.ShowPct = show
}

// SetLabel sets the label for the meter
func (m *Meter) SetLabel(label string) {
	m.Label = label
}

// drawTextStyled draws a string at the specified position with the given style
func (m *Meter) drawTextStyled(x, y int, text string, style tcell.Style) {
	for i, r := range text {
		m.Screen.SetContent(x+i, y, r, nil, style)
	}
}

// interpolateColor interpolates between two colors based on position (0-1)
func interpolateColor(start, end tcell.Color, position float64) tcell.Color {
	r1, g1, b1 := start.RGB()
	r2, g2, b2 := end.RGB()

	r := uint8(float64(r1) + position*(float64(r2)-float64(r1)))
	g := uint8(float64(g1) + position*(float64(g2)-float64(g1)))
	b := uint8(float64(b1) + position*(float64(b2)-float64(b1)))

	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}
