package draw

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestDrawHLine(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	DrawHLine(screen, 0, 0, 10, style)

	// Check first and last cells
	mainc, _, _, _ := screen.GetContent(0, 0)
	if mainc != '─' {
		t.Errorf("Expected horizontal line character at start, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(9, 0)
	if mainc != '─' {
		t.Errorf("Expected horizontal line character at end, got %c", mainc)
	}
}

func TestDrawVLine(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	DrawVLine(screen, 0, 0, 10, style)

	// Check first and last cells
	mainc, _, _, _ := screen.GetContent(0, 0)
	if mainc != '│' {
		t.Errorf("Expected vertical line character at start, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(0, 9)
	if mainc != '│' {
		t.Errorf("Expected vertical line character at end, got %c", mainc)
	}
}

func TestDrawBox(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	DrawBox(screen, 0, 0, 10, 5, style)

	// Check corners
	corners := []struct {
		x, y int
		char rune
	}{
		{0, 0, '┌'},   // Top-left
		{9, 0, '┐'},   // Top-right
		{0, 4, '└'},   // Bottom-left
		{9, 4, '┘'},   // Bottom-right
	}

	for _, c := range corners {
		mainc, _, _, _ := screen.GetContent(c.x, c.y)
		if mainc != c.char {
			t.Errorf("Expected corner character %c at (%d,%d), got %c", c.char, c.x, c.y, mainc)
		}
	}
}

func TestDrawText(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Hello"
	DrawText(screen, 0, 0, style, text)

	// Check each character
	for i, char := range text {
		mainc, _, _, _ := screen.GetContent(i, 0)
		if mainc != char {
			t.Errorf("Expected character %c at position %d, got %c", char, i, mainc)
		}
	}
}

func TestDrawTextCentered(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Center"
	width := 10
	DrawTextCentered(screen, 0, 0, width, style, text)

	// Calculate expected start position
	startX := (width - len(text)) / 2

	// Check each character
	for i, char := range text {
		mainc, _, _, _ := screen.GetContent(startX+i, 0)
		if mainc != char {
			t.Errorf("Expected character %c at position %d, got %c", char, startX+i, mainc)
		}
	}
}

func TestDrawTextRight(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Right"
	width := 10
	DrawTextRight(screen, 0, 0, width, style, text)

	// Calculate expected start position
	startX := width - len(text)

	// Check each character
	for i, char := range text {
		mainc, _, _, _ := screen.GetContent(startX+i, 0)
		if mainc != char {
			t.Errorf("Expected character %c at position %d, got %c", char, startX+i, mainc)
		}
	}
}
