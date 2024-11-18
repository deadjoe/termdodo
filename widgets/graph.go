package widgets

import (
	"github.com/deadjoe/termdodo/symbols"
	"github.com/deadjoe/termdodo/theme"
	tcell "github.com/gdamore/tcell/v2"
)

// GraphStyle represents the style of graph to be drawn
type GraphStyle int

// Graph styles
const (
	// GraphStyleBraille uses Braille characters for drawing the graph
	GraphStyleBraille GraphStyle = iota
	// GraphStyleBlock uses block characters for drawing the graph
	GraphStyleBlock
	// GraphStyleTTY uses TTY characters for drawing the graph
	GraphStyleTTY
)

// Graph represents a percentage graph widget
type Graph struct {
	X, Y          int
	Width, Height int
	Style         tcell.Style
	Screen        tcell.Screen
	GraphStyle    GraphStyle
	Data          []float64
	MaxValue      float64
	MinValue      float64
	Inverted      bool
}

// NewGraph creates a new graph widget
func NewGraph(screen tcell.Screen, x, y, width, height int) *Graph {
	return &Graph{
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Style:    theme.Current.GetStyle(),
		Screen:   screen,
		Data:     make([]float64, 0),
		MaxValue: 100,
		MinValue: 0,
		Inverted: false,
	}
}

// Draw draws the graph on the screen
func (g *Graph) Draw() {
	if len(g.Data) == 0 {
		return
	}

	// Calculate the scale factor
	scale := float64(g.Height) / (g.MaxValue - g.MinValue)

	// Get the pattern set based on graph style
	var patterns []string
	switch g.GraphStyle {
	case GraphStyleBraille:
		patterns = symbols.BraillePatterns
	case GraphStyleBlock:
		patterns = symbols.BlockPatterns
	case GraphStyleTTY:
		patterns = symbols.TTYPatterns
	}

	// Draw each data point
	for i, value := range g.Data {
		if i >= g.Width {
			break
		}

		// Calculate the height of this column
		height := int((value - g.MinValue) * scale)
		if g.Inverted {
			height = g.Height - height
		}

		// Calculate the color position (0.0 - 1.0)
		position := float64(height) / float64(g.Height)

		// Get the color for this position
		style := theme.Current.GetGradientStyle(position)

		// Draw the column
		x := g.X + i
		for y := 0; y < g.Height; y++ {
			var pattern string
			if y < height {
				pattern = patterns[len(patterns)-1] // Full block
			} else if y == height {
				// Calculate partial block
				fraction := (value - float64(height)/scale) * scale
				patternIndex := int(fraction * float64(len(patterns)-1))
				pattern = patterns[patternIndex]
			} else {
				pattern = patterns[0] // Empty block
			}

			g.drawTextStyled(x, g.Y+y, pattern, style)
		}
	}
}

// SetData sets the data points for the graph
func (g *Graph) SetData(data []float64) {
	g.Data = data
}

// SetStyle sets the style for the graph
func (g *Graph) SetStyle(style tcell.Style) {
	g.Style = style
}

// SetGraphStyle sets the style of graph to be drawn
func (g *Graph) SetGraphStyle(style GraphStyle) {
	g.GraphStyle = style
}

// drawTextStyled draws a string at the specified position with the given style
func (g *Graph) drawTextStyled(x, y int, text string, style tcell.Style) {
	for i, r := range text {
		g.Screen.SetContent(x+i, y, r, nil, style)
	}
}

// Clear clears the graph data
func (g *Graph) Clear() {
	g.Data = make([]float64, 0)
}

// SetRange sets the value range for the graph
func (g *Graph) SetRange(min, max float64) {
	g.MinValue = min
	g.MaxValue = max
}

// SetInverted sets whether to invert the graph
func (g *Graph) SetInverted(inverted bool) {
	g.Inverted = inverted
}
