package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewStatusBar(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	bar := NewStatusBar(screen, 0, 0, 80)
	if bar == nil {
		t.Fatal("Expected non-nil StatusBar")
	}

	if bar.Width != 80 {
		t.Errorf("Expected width to be 80, got %d", bar.Width)
	}
}

func TestStatusBarSetItems(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	bar := NewStatusBar(screen, 0, 0, 80)
	items := []StatusItem{
		{Text: "Connected", Style: tcell.StyleDefault},
		{Text: "CPU: 45%", Style: tcell.StyleDefault.Foreground(tcell.ColorGreen)},
		{Text: "Memory: 2.5GB", Style: tcell.StyleDefault.Foreground(tcell.ColorYellow)},
	}

	bar.SetItems(items)

	if len(bar.Items) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(bar.Items))
	}

	for i, item := range bar.Items {
		if item.Text != items[i].Text {
			t.Errorf("Item %d: expected text %q, got %q", i, items[i].Text, item.Text)
		}
	}
}

func TestStatusBarUpdateItem(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	bar := NewStatusBar(screen, 0, 0, 80)
	items := []StatusItem{
		{Text: "Connected", Style: tcell.StyleDefault},
		{Text: "CPU: 45%", Style: tcell.StyleDefault},
	}
	bar.SetItems(items)

	// Test updating existing item
	newStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	bar.UpdateItem(0, "Disconnected", newStyle)
	
	if bar.Items[0].Text != "Disconnected" {
		t.Errorf("Expected updated text to be 'Disconnected', got %q", bar.Items[0].Text)
	}

	// Test updating out of bounds index (should not panic)
	bar.UpdateItem(99, "Invalid", tcell.StyleDefault)
}

func TestStatusBarClear(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	bar := NewStatusBar(screen, 0, 0, 80)
	items := []StatusItem{
		{Text: "Connected", Style: tcell.StyleDefault},
		{Text: "CPU: 45%", Style: tcell.StyleDefault},
	}
	bar.SetItems(items)

	bar.Clear()

	if len(bar.Items) != 0 {
		t.Errorf("Expected empty items after clear, got %d items", len(bar.Items))
	}
}

func TestStatusBarSetSeparator(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	bar := NewStatusBar(screen, 0, 0, 80)
	testSep := " | "
	bar.SetSeparator(testSep)

	if bar.Separator != testSep {
		t.Errorf("Expected separator to be %q, got %q", testSep, bar.Separator)
	}
}
