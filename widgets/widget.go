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
