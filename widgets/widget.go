package widgets

import "github.com/gdamore/tcell/v2"

// Widget defines the interface that all widgets must implement
type Widget interface {
	// Draw renders the widget on the screen
	Draw()

	// Clear removes the widget from the screen
	Clear()

	// GetBounds returns the widget's position and size
	GetBounds() (x, y, width, height int)

	// SetBounds sets the widget's position and size
	SetBounds(x, y, width, height int)

	// SetStyle sets the widget's style
	SetStyle(style tcell.Style)
}

// BaseWidget provides common functionality for all widgets
type BaseWidget struct {
	X, Y          int
	Width, Height int
	Style         tcell.Style
	Screen        tcell.Screen
}

// NewBaseWidget creates a new base widget
func NewBaseWidget(screen tcell.Screen, x, y, width, height int) BaseWidget {
	return BaseWidget{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Screen: screen,
	}
}

// GetBounds returns the widget's position and size
func (w *BaseWidget) GetBounds() (x, y, width, height int) {
	return w.X, w.Y, w.Width, w.Height
}

// SetBounds sets the widget's position and size
func (w *BaseWidget) SetBounds(x, y, width, height int) {
	w.X = x
	w.Y = y
	w.Width = width
	w.Height = height
}

// SetStyle sets the widget's style
func (w *BaseWidget) SetStyle(style tcell.Style) {
	w.Style = style
}

// Clear removes the widget from the screen
func (w *BaseWidget) Clear() {
	for y := w.Y; y < w.Y+w.Height; y++ {
		for x := w.X; x < w.X+w.Width; x++ {
			w.Screen.SetContent(x, y, ' ', nil, w.Style)
		}
	}
}

// DrawBorder draws a border around the widget
func DrawBorder(screen tcell.Screen, x, y, width, height int, style tcell.Style) {
	// Draw corners
	screen.SetContent(x, y, '┌', nil, style)
	screen.SetContent(x+width-1, y, '┐', nil, style)
	screen.SetContent(x, y+height-1, '└', nil, style)
	screen.SetContent(x+width-1, y+height-1, '┘', nil, style)

	// Draw horizontal lines
	for i := x + 1; i < x+width-1; i++ {
		screen.SetContent(i, y, '─', nil, style)
		screen.SetContent(i, y+height-1, '─', nil, style)
	}

	// Draw vertical lines
	for i := y + 1; i < y+height-1; i++ {
		screen.SetContent(x, i, '│', nil, style)
		screen.SetContent(x+width-1, i, '│', nil, style)
	}
}
