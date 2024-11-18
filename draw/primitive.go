package draw

import (
	"github.com/gdamore/tcell/v2"
)

// HLine draws a horizontal line
func HLine(screen tcell.Screen, x, y, width int, style tcell.Style) {
	for i := 0; i < width; i++ {
		screen.SetContent(x+i, y, '─', nil, style)
	}
}

// VLine draws a vertical line
func VLine(screen tcell.Screen, x, y, height int, style tcell.Style) {
	for i := 0; i < height; i++ {
		screen.SetContent(x, y+i, '│', nil, style)
	}
}

// Rect draws a rectangle with the specified dimensions
func Rect(screen tcell.Screen, x, y, width, height int, style tcell.Style) {
	// Draw corners
	screen.SetContent(x, y, '┌', nil, style)
	screen.SetContent(x+width-1, y, '┐', nil, style)
	screen.SetContent(x, y+height-1, '└', nil, style)
	screen.SetContent(x+width-1, y+height-1, '┘', nil, style)

	// Draw borders
	HLine(screen, x+1, y, width-2, style)
	HLine(screen, x+1, y+height-1, width-2, style)
	VLine(screen, x, y+1, height-2, style)
	VLine(screen, x+width-1, y+1, height-2, style)
}

// Text draws text at the specified position
func Text(screen tcell.Screen, x, y int, style tcell.Style, text string) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

// TextCentered draws text centered at the specified position
func TextCentered(screen tcell.Screen, x, y, width int, style tcell.Style, text string) {
	if width <= 0 {
		return
	}

	textWidth := len(text)
	if textWidth > width {
		text = text[:width]
		textWidth = width
	}

	startX := x + (width-textWidth)/2
	Text(screen, startX, y, style, text)
}

// TextRight draws text aligned to the right at the specified position
func TextRight(screen tcell.Screen, x, y, width int, style tcell.Style, text string) {
	if width <= 0 {
		return
	}

	textWidth := len(text)
	if textWidth > width {
		text = text[:width]
		textWidth = width
	}

	startX := x + width - textWidth
	Text(screen, startX, y, style, text)
}
