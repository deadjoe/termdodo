package widgets

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell/v2"
	"termdodo/symbols"
	"termdodo/theme"
)

// Meter represents a percentage meter widget
type Meter struct {
	X, Y    int
	Width   int
	Style   tcell.Style
	Screen  tcell.Screen
	Value   float64
	ShowPct bool
	
	// New fields for block style and gradient
	BlockStyle    bool
	BlockSpacing  int
	StartColor    tcell.Color
	EndColor      tcell.Color
	UseGradient   bool
}

// NewMeter creates a new meter widget
func NewMeter(screen tcell.Screen, x, y, width int) *Meter {
	mainFg := theme.GetColor(theme.Current.MainFg)
	return &Meter{
		X:            x,
		Y:            y,
		Width:        width,
		Style:        theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		Screen:       screen,
		Value:        0,
		ShowPct:      true,
		BlockStyle:   false,
		BlockSpacing: 0,
		StartColor:   mainFg,
		EndColor:     mainFg,
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
	// Calculate filled width (accounting for block spacing if enabled)
	totalWidth := m.Width
	if m.BlockStyle {
		// Adjust width to account for spacing
		blockCount := m.Width / (1 + m.BlockSpacing)
		if blockCount < 1 {
			blockCount = 1
		}
		totalWidth = blockCount
	}

	filledWidth := int(float64(totalWidth) * m.Value)
	if filledWidth > totalWidth {
		filledWidth = totalWidth
	}

	// Draw meter blocks
	pos := m.X
	for i := 0; i < totalWidth; i++ {
		var style tcell.Style
		if i < filledWidth {
			if m.UseGradient {
				// Calculate position in the gradient (0.0 - 1.0)
				position := float64(i) / float64(totalWidth)
				r1, g1, b1 := m.StartColor.RGB()
				r2, g2, b2 := m.EndColor.RGB()
				
				// Linear interpolation between colors
				r := uint8(float64(r1) + position*float64(r2-r1))
				g := uint8(float64(g1) + position*float64(g2-g1))
				b := uint8(float64(b1) + position*float64(b2-b1))
				
				color := tcell.NewRGBColor(int32(r), int32(g), int32(b))
				style = tcell.StyleDefault.Foreground(color)
			} else {
				style = m.Style
			}
		} else {
			style = theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg)
		}

		if m.BlockStyle {
			// 使用密集点字符组成紧凑的正方形
			m.drawTextStyled(pos, m.Y, "⠿", style)      // 使用布莱叶密集点字符
			pos += 1 + m.BlockSpacing  // 每个方块占一个字符宽度
		} else {
			m.drawTextStyled(pos, m.Y, symbols.Meter, style)
			pos++
		}
	}

	// Draw percentage if enabled
	if m.ShowPct {
		pct := fmt.Sprintf(" %3.0f%%", m.Value*100)
		style := theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg)
		for i, ch := range pct {
			m.drawTextStyled(pos+i, m.Y, string(ch), style)
		}
	}
}

// SetValue sets the current value of the meter (0-100)
func (m *Meter) SetValue(value float64) {
	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	m.Value = value / 100
}

// SetStyle sets the style for the meter
func (m *Meter) SetStyle(style tcell.Style) {
	m.Style = style
}

// SetShowPercentage sets whether to show the percentage value
func (m *Meter) SetShowPercentage(show bool) {
	m.ShowPct = show
}

// drawTextStyled draws a string at the specified position with the given style
func (m *Meter) drawTextStyled(x, y int, text string, style tcell.Style) {
	for i, r := range text {
		m.Screen.SetContent(x+i, y, r, nil, style)
	}
}
