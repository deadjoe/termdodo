package draw

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestHLine(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	HLine(screen, 0, 0, 10, style)

	mainc, _, _, _ := screen.GetContent(0, 0)
	if mainc != '─' {
		t.Errorf("Expected horizontal line character at start, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(9, 0)
	if mainc != '─' {
		t.Errorf("Expected horizontal line character at end, got %c", mainc)
	}
}

func TestVLine(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	VLine(screen, 0, 0, 10, style)

	mainc, _, _, _ := screen.GetContent(0, 0)
	if mainc != '│' {
		t.Errorf("Expected vertical line character at start, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(0, 9)
	if mainc != '│' {
		t.Errorf("Expected vertical line character at end, got %c", mainc)
	}
}

func TestRect(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	Rect(screen, 0, 0, 10, 5, style)

	// Check corners
	mainc, _, _, _ := screen.GetContent(0, 0)
	if mainc != '┌' {
		t.Errorf("Expected top-left corner, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(9, 0)
	if mainc != '┐' {
		t.Errorf("Expected top-right corner, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(0, 4)
	if mainc != '└' {
		t.Errorf("Expected bottom-left corner, got %c", mainc)
	}

	mainc, _, _, _ = screen.GetContent(9, 4)
	if mainc != '┘' {
		t.Errorf("Expected bottom-right corner, got %c", mainc)
	}
}

func TestText(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Hello"
	Text(screen, 0, 0, style, text)

	// Check each character
	for i, expected := range text {
		mainc, _, _, _ := screen.GetContent(i, 0)
		if mainc != expected {
			t.Errorf("Expected %c at position %d, got %c", expected, i, mainc)
		}
	}
}

func TestTextCentered(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Center"
	width := 10
	TextCentered(screen, 0, 0, width, style, text)

	// Check each character
	startX := (width - len(text)) / 2
	for i, expected := range text {
		mainc, _, _, _ := screen.GetContent(startX+i, 0)
		if mainc != expected {
			t.Errorf("Expected %c at position %d, got %c", expected, i, mainc)
		}
	}
}

func TestTextRight(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	style := tcell.StyleDefault
	text := "Right"
	width := 10
	TextRight(screen, 0, 0, width, style, text)

	// Check each character
	startX := width - len(text)
	for i, expected := range text {
		mainc, _, _, _ := screen.GetContent(startX+i, 0)
		if mainc != expected {
			t.Errorf("Expected %c at position %d, got %c", expected, i, mainc)
		}
	}
}
