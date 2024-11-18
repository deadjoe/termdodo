package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewBaseWidget(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	x, y := 5, 10
	width, height := 20, 15
	widget := NewBaseWidget(screen, x, y, width, height)

	// Check initial values
	if widget.X != x || widget.Y != y {
		t.Errorf("Expected position (%d,%d), got (%d,%d)", x, y, widget.X, widget.Y)
	}
	if widget.Width != width || widget.Height != height {
		t.Errorf("Expected size (%d,%d), got (%d,%d)", width, height, widget.Width, widget.Height)
	}
	if widget.Screen != screen {
		t.Error("Screen not properly set")
	}
	if widget.Style != tcell.StyleDefault {
		t.Error("Style should default to tcell.StyleDefault")
	}
}

func TestBaseWidgetGetBounds(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	x, y := 5, 10
	width, height := 20, 15
	widget := NewBaseWidget(screen, x, y, width, height)

	gotX, gotY, gotWidth, gotHeight := widget.GetBounds()
	if gotX != x || gotY != y || gotWidth != width || gotHeight != height {
		t.Errorf("GetBounds() = (%d,%d,%d,%d), want (%d,%d,%d,%d)",
			gotX, gotY, gotWidth, gotHeight, x, y, width, height)
	}
}

func TestBaseWidgetSetBounds(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	widget := NewBaseWidget(screen, 0, 0, 10, 10)
	newX, newY := 15, 20
	newWidth, newHeight := 30, 25

	widget.SetBounds(newX, newY, newWidth, newHeight)
	gotX, gotY, gotWidth, gotHeight := widget.GetBounds()

	if gotX != newX || gotY != newY || gotWidth != newWidth || gotHeight != newHeight {
		t.Errorf("After SetBounds(%d,%d,%d,%d), got (%d,%d,%d,%d)",
			newX, newY, newWidth, newHeight, gotX, gotY, gotWidth, gotHeight)
	}
}

func TestBaseWidgetSetStyle(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	widget := NewBaseWidget(screen, 0, 0, 10, 10)
	newStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow)

	widget.SetStyle(newStyle)
	if widget.Style != newStyle {
		t.Error("Style not properly set")
	}
}

func TestBaseWidgetClear(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	widget := NewBaseWidget(screen, 5, 5, 3, 2)
	style := tcell.StyleDefault.Background(tcell.ColorBlue)
	widget.SetStyle(style)

	// First fill the area with some content
	for y := widget.Y; y < widget.Y+widget.Height; y++ {
		for x := widget.X; x < widget.X+widget.Width; x++ {
			screen.SetContent(x, y, 'X', nil, style)
		}
	}

	// Clear the widget
	widget.Clear()

	// Check that all cells in the widget area are cleared
	for y := widget.Y; y < widget.Y+widget.Height; y++ {
		for x := widget.X; x < widget.X+widget.Width; x++ {
			mainc, _, _, _ := screen.GetContent(x, y)
			if mainc != ' ' {
				t.Errorf("Expected space at (%d,%d), got %c", x, y, mainc)
			}
		}
	}
}

func TestDrawBorder(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	screen.SetSize(100, 100)

	x, y := 5, 5
	width, height := 4, 3
	style := tcell.StyleDefault

	DrawBorder(screen, x, y, width, height, style)

	// Test corners
	corners := []struct {
		x, y int
		char rune
	}{
		{x, y, '┌'},                          // Top-left
		{x + width - 1, y, '┐'},              // Top-right
		{x, y + height - 1, '└'},             // Bottom-left
		{x + width - 1, y + height - 1, '┘'}, // Bottom-right
	}

	for _, c := range corners {
		mainc, _, _, _ := screen.GetContent(c.x, c.y)
		if mainc != c.char {
			t.Errorf("At (%d,%d): expected %c, got %c", c.x, c.y, c.char, mainc)
		}
	}

	// Test horizontal lines
	for i := x + 1; i < x+width-1; i++ {
		// Top line
		mainc, _, _, _ := screen.GetContent(i, y)
		if mainc != '─' {
			t.Errorf("At (%d,%d): expected ─, got %c", i, y, mainc)
		}
		// Bottom line
		mainc, _, _, _ = screen.GetContent(i, y+height-1)
		if mainc != '─' {
			t.Errorf("At (%d,%d): expected ─, got %c", i, y+height-1, mainc)
		}
	}

	// Test vertical lines
	for i := y + 1; i < y+height-1; i++ {
		// Left line
		mainc, _, _, _ := screen.GetContent(x, i)
		if mainc != '│' {
			t.Errorf("At (%d,%d): expected │, got %c", x, i, mainc)
		}
		// Right line
		mainc, _, _, _ = screen.GetContent(x+width-1, i)
		if mainc != '│' {
			t.Errorf("At (%d,%d): expected │, got %c", x+width-1, i, mainc)
		}
	}
}
