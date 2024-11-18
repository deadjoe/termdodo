package widgets

import (
	"fmt"

	"github.com/deadjoe/termdodo/theme"
	"github.com/gdamore/tcell/v2"
)

// MeterItem represents a single meter in the multi meter widget
type MeterItem struct {
	Label          string
	Value          float64
	MaxValue       float64
	Style          tcell.Style
	GradientColors []tcell.Color
}

// MultiMeter represents a multi meter widget
type MultiMeter struct {
	X, Y          int
	Width, Height int
	Screen        tcell.Screen
	Style         tcell.Style
	LabelStyle    tcell.Style

	Items       []MeterItem
	ShowLabels  bool
	ShowValues  bool
	ShowBorder  bool
	Orientation Orientation
	LabelWidth  int
	MeterHeight int
	Spacing     int
}

// Orientation represents the orientation of the multi meter
type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

// NewMultiMeter creates a new multi meter widget
func NewMultiMeter(screen tcell.Screen, x, y, width, height int) *MultiMeter {
	return &MultiMeter{
		X:           x,
		Y:           y,
		Width:       width,
		Height:      height,
		Screen:      screen,
		Style:       theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		LabelStyle:  theme.GetStyle(theme.Current.Title, theme.Current.MainBg),
		ShowLabels:  true,
		ShowValues:  true,
		ShowBorder:  true,
		Orientation: Vertical,
		LabelWidth:  20,
		MeterHeight: 1,
		Spacing:     1,
	}
}

// SetItems sets the meter items
func (m *MultiMeter) SetItems(items []MeterItem) {
	m.Items = items
}

// AddItem adds a meter item
func (m *MultiMeter) AddItem(item MeterItem) {
	m.Items = append(m.Items, item)
}

// ClearItems clears all meter items
func (m *MultiMeter) ClearItems() {
	m.Items = nil
}

// Draw draws the multi meter
func (m *MultiMeter) Draw() {
	if len(m.Items) == 0 {
		return
	}

	if m.ShowBorder {
		m.drawBorder()
	}

	startX := m.X
	startY := m.Y
	if m.ShowBorder {
		startX++
		startY++
	}

	switch m.Orientation {
	case Vertical:
		m.drawVertical(startX, startY)
	case Horizontal:
		m.drawHorizontal(startX, startY)
	}
}

// drawBorder draws the widget border
func (m *MultiMeter) drawBorder() {
	style := m.Style

	// Draw corners
	m.Screen.SetContent(m.X, m.Y, '┌', nil, style)
	m.Screen.SetContent(m.X+m.Width-1, m.Y, '┐', nil, style)
	m.Screen.SetContent(m.X, m.Y+m.Height-1, '└', nil, style)
	m.Screen.SetContent(m.X+m.Width-1, m.Y+m.Height-1, '┘', nil, style)

	// Draw horizontal lines
	for x := m.X + 1; x < m.X+m.Width-1; x++ {
		m.Screen.SetContent(x, m.Y, '─', nil, style)
		m.Screen.SetContent(x, m.Y+m.Height-1, '─', nil, style)
	}

	// Draw vertical lines
	for y := m.Y + 1; y < m.Y+m.Height-1; y++ {
		m.Screen.SetContent(m.X, y, '│', nil, style)
		m.Screen.SetContent(m.X+m.Width-1, y, '│', nil, style)
	}
}

// drawLabel draws a label for a meter item
func (m *MultiMeter) drawLabel(x, y int, item MeterItem) {
	labelStyle := m.LabelStyle
	if item.Style != (tcell.Style{}) {
		labelStyle = item.Style
	}
	for i, r := range item.Label {
		if i >= m.LabelWidth {
			break
		}
		m.Screen.SetContent(x+i, y, r, nil, labelStyle)
	}
}

// drawMeterBar draws a meter bar for a meter item
func (m *MultiMeter) drawMeterBar(x, y, width int, item MeterItem) {
	// Calculate value percentage
	percentage := item.Value / item.MaxValue
	if percentage > 1.0 {
		percentage = 1.0
	}

	// Calculate filled width
	filledWidth := int(float64(width) * percentage)
	if filledWidth > width {
		filledWidth = width
	}

	// Draw filled part
	style := m.Style
	if item.Style != (tcell.Style{}) {
		style = item.Style
	}
	for i := 0; i < filledWidth; i++ {
		m.Screen.SetContent(x+i, y, '█', nil, style)
	}

	// Draw empty part
	emptyStyle := style.Background(tcell.ColorBlack)
	for i := filledWidth; i < width; i++ {
		m.Screen.SetContent(x+i, y, '░', nil, emptyStyle)
	}
}

// drawValue draws a value for a meter item
func (m *MultiMeter) drawValue(x, y int, item MeterItem) {
	valueStr := formatValue(item.Value)
	style := m.Style
	if item.Style != (tcell.Style{}) {
		style = item.Style
	}
	for i, r := range valueStr {
		m.Screen.SetContent(x+i, y, r, nil, style)
	}
}

// drawVertical draws meters vertically
func (m *MultiMeter) drawVertical(startX, startY int) {
	availableHeight := m.Height
	if m.ShowBorder {
		availableHeight -= 2
	}

	itemHeight := m.MeterHeight
	if m.ShowLabels {
		itemHeight++
	}
	if m.ShowValues {
		itemHeight++
	}
	itemHeight += m.Spacing

	y := startY
	for _, item := range m.Items {
		if y+itemHeight > startY+availableHeight {
			break
		}

		// Draw label
		if m.ShowLabels {
			m.drawLabel(startX, y, item)
			y++
		}

		// Draw meter
		m.drawMeterBar(startX, y, m.Width-2, item)
		y++

		// Draw value
		if m.ShowValues {
			m.drawValue(startX, y, item)
			y++
		}

		y += m.Spacing
	}
}

// drawHorizontal draws meters horizontally
func (m *MultiMeter) drawHorizontal(startX, startY int) {
	availableWidth := m.Width
	if m.ShowBorder {
		availableWidth -= 2
	}

	// Calculate item width
	itemWidth := availableWidth / len(m.Items)
	if itemWidth < 1 {
		itemWidth = 1
	}

	x := startX
	for _, item := range m.Items {
		if x+itemWidth > startX+availableWidth {
			break
		}

		y := startY
		labelWidth := itemWidth

		// Draw label
		if m.ShowLabels {
			m.drawLabel(x, y, item)
			y++
		}

		// Draw meter
		m.drawMeterBar(x, y, labelWidth, item)
		y++

		// Draw value
		if m.ShowValues {
			m.drawValue(x, y, item)
		}

		x += itemWidth + m.Spacing
	}
}

// formatValue formats a float value as a string
func formatValue(value float64) string {
	if value >= 100 {
		return fmt.Sprintf("%.0f%%", value)
	}
	return fmt.Sprintf("%.1f%%", value)
}

// truncateString truncates a string to the specified width
func truncateString(s string, width int) string {
	if len(s) <= width {
		return s
	}
	return s[:width-3] + "..."
}

// SetShowLabels sets whether to show labels
func (m *MultiMeter) SetShowLabels(show bool) {
	m.ShowLabels = show
}

// SetShowValues sets whether to show values
func (m *MultiMeter) SetShowValues(show bool) {
	m.ShowValues = show
}

// SetShowBorder sets whether to show the border
func (m *MultiMeter) SetShowBorder(show bool) {
	m.ShowBorder = show
}

// SetOrientation sets the orientation of the meters
func (m *MultiMeter) SetOrientation(orientation Orientation) {
	m.Orientation = orientation
}

// SetLabelWidth sets the width of labels
func (m *MultiMeter) SetLabelWidth(width int) {
	m.LabelWidth = width
}

// SetMeterHeight sets the height of meters
func (m *MultiMeter) SetMeterHeight(height int) {
	m.MeterHeight = height
}

// SetSpacing sets the spacing between meters
func (m *MultiMeter) SetSpacing(spacing int) {
	m.Spacing = spacing
}

// GetHeight returns the total height of the widget
func (m *MultiMeter) GetHeight() int {
	return m.Height
}

// GetWidth returns the total width of the widget
func (m *MultiMeter) GetWidth() int {
	return m.Width
}

// SetStyle sets the default style for the widget
func (m *MultiMeter) SetStyle(style tcell.Style) {
	m.Style = style
}

// SetLabelStyle sets the style for labels
func (m *MultiMeter) SetLabelStyle(style tcell.Style) {
	m.LabelStyle = style
}
