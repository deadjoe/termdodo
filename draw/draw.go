package draw

import (
	"github.com/deadjoe/termdodo/symbols"
	"github.com/gdamore/tcell/v2"
)

// DrawHLine draws a horizontal line
func DrawHLine(screen tcell.Screen, x, y, width int, style tcell.Style) {
	for i := 0; i < width; i++ {
		screen.SetContent(x+i, y, symbols.HLine, nil, style)
	}
}

// DrawVLine draws a vertical line
func DrawVLine(screen tcell.Screen, x, y, height int, style tcell.Style) {
	for i := 0; i < height; i++ {
		screen.SetContent(x, y+i, symbols.VLine, nil, style)
	}
}

// DrawBox draws a box with the given dimensions
func DrawBox(screen tcell.Screen, x, y, width, height int, style tcell.Style) {
	// Draw corners
	screen.SetContent(x, y, symbols.TLCorner, nil, style)
	screen.SetContent(x+width-1, y, symbols.TRCorner, nil, style)
	screen.SetContent(x, y+height-1, symbols.BLCorner, nil, style)
	screen.SetContent(x+width-1, y+height-1, symbols.BRCorner, nil, style)

	// Draw horizontal lines
	for i := 1; i < width-1; i++ {
		screen.SetContent(x+i, y, symbols.HLine, nil, style)
		screen.SetContent(x+i, y+height-1, symbols.HLine, nil, style)
	}

	// Draw vertical lines
	for i := 1; i < height-1; i++ {
		screen.SetContent(x, y+i, symbols.VLine, nil, style)
		screen.SetContent(x+width-1, y+i, symbols.VLine, nil, style)
	}
}

// DrawText draws text at the specified position
func DrawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

// DrawTextCentered draws centered text at the specified position
func DrawTextCentered(screen tcell.Screen, x, y, width int, text string, style tcell.Style) {
	textWidth := len(text)
	if textWidth > width {
		text = text[:width]
		textWidth = width
	}
	startX := x + (width-textWidth)/2
	DrawText(screen, startX, y, text, style)
}

// DrawTextRight draws right-aligned text at the specified position
func DrawTextRight(screen tcell.Screen, x, y, width int, text string, style tcell.Style) {
	textWidth := len(text)
	if textWidth > width {
		text = text[textWidth-width:]
		textWidth = width
	}
	startX := x + width - textWidth
	DrawText(screen, startX, y, text, style)
}
