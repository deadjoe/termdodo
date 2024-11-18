package draw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/deadjoe/termdodo/symbols"
	"github.com/deadjoe/termdodo/theme"
)

// Box represents a box widget with borders
type Box struct {
	X, Y          int
	Width, Height int
	Title         string
	Round         bool
	Screen        tcell.Screen
	Style         tcell.Style
}

// NewBox creates a new box widget
func NewBox(screen tcell.Screen, x, y, width, height int) *Box {
	return &Box{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Screen: screen,
		Style:  theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
	}
}

// SetTitle sets the box title
func (b *Box) SetTitle(title string) {
	b.Title = title
}

// SetRound sets whether to use rounded corners
func (b *Box) SetRound(round bool) {
	b.Round = round
}

// SetStyle sets the box style
func (b *Box) SetStyle(style tcell.Style) {
	b.Style = style
}

// InnerX returns the inner X coordinate
func (b *Box) InnerX() int {
	return b.X + 1
}

// InnerY returns the inner Y coordinate
func (b *Box) InnerY() int {
	return b.Y + 1
}

// InnerWidth returns the inner width
func (b *Box) InnerWidth() int {
	return b.Width - 2
}

// InnerHeight returns the inner height
func (b *Box) InnerHeight() int {
	return b.Height - 2
}

// Draw draws the box on the screen
func (b *Box) Draw() {
	// Get box characters based on style
	var topLeft, topRight, bottomLeft, bottomRight, horizontal, vertical string
	if b.Round {
		topLeft = symbols.BoxDrawingRoundTopLeft
		topRight = symbols.BoxDrawingRoundTopRight
		bottomLeft = symbols.BoxDrawingRoundBottomLeft
		bottomRight = symbols.BoxDrawingRoundBottomRight
	} else {
		topLeft = symbols.BoxDrawingTopLeft
		topRight = symbols.BoxDrawingTopRight
		bottomLeft = symbols.BoxDrawingBottomLeft
		bottomRight = symbols.BoxDrawingBottomRight
	}
	horizontal = symbols.BoxDrawingHorizontal
	vertical = symbols.BoxDrawingVertical

	// Draw corners
	b.drawRune(b.X, b.Y, []rune(topLeft)[0])
	b.drawRune(b.X+b.Width-1, b.Y, []rune(topRight)[0])
	b.drawRune(b.X, b.Y+b.Height-1, []rune(bottomLeft)[0])
	b.drawRune(b.X+b.Width-1, b.Y+b.Height-1, []rune(bottomRight)[0])

	// Draw horizontal borders
	for x := b.X + 1; x < b.X+b.Width-1; x++ {
		b.drawRune(x, b.Y, []rune(horizontal)[0])
		b.drawRune(x, b.Y+b.Height-1, []rune(horizontal)[0])
	}

	// Draw vertical borders
	for y := b.Y + 1; y < b.Y+b.Height-1; y++ {
		b.drawRune(b.X, y, []rune(vertical)[0])
		b.drawRune(b.X+b.Width-1, y, []rune(vertical)[0])
	}

	// Draw title if set
	if b.Title != "" {
		titleStyle := theme.GetStyle(theme.Current.Title, theme.Current.MainBg)
		titleX := b.X + (b.Width-len(b.Title))/2
		for i, r := range b.Title {
			b.Screen.SetContent(titleX+i, b.Y, r, nil, titleStyle)
		}
	}
}

// drawRune draws a single rune at the specified position
func (b *Box) drawRune(x, y int, r rune) {
	b.Screen.SetContent(x, y, r, nil, b.Style)
}
