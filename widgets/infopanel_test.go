package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewInfoPanel(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	if panel == nil {
		t.Fatal("Expected non-nil InfoPanel")
	}

	if panel.Width != 80 || panel.Height != 10 {
		t.Errorf("Expected panel dimensions to be 80x10, got %dx%d", panel.Width, panel.Height)
	}
}

func TestInfoPanelSetTitle(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	testTitle := "System Info"
	panel.SetTitle(testTitle)

	if panel.Title != testTitle {
		t.Errorf("Expected title to be %q, got %q", testTitle, panel.Title)
	}
}

func TestInfoPanelSetItems(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	items := []InfoItem{
		{Label: "CPU", Value: "45%"},
		{Label: "Memory", Value: "2.5GB"},
		{Label: "Disk", Value: "120GB"},
	}

	panel.SetItems(items)

	if len(panel.Items) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(panel.Items))
	}

	for i, item := range panel.Items {
		if item.Label != items[i].Label {
			t.Errorf("Item %d: expected label %q, got %q", i, items[i].Label, item.Label)
		}
		if item.Value != items[i].Value {
			t.Errorf("Item %d: expected value %q, got %q", i, items[i].Value, item.Value)
		}
	}
}

func TestInfoPanelUpdateItem(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	items := []InfoItem{
		{Label: "CPU", Value: "45%"},
		{Label: "Memory", Value: "2.5GB"},
	}
	panel.SetItems(items)

	// Test updating existing item
	panel.UpdateItem("CPU", "55%")
	if panel.Items[0].Value != "55%" {
		t.Errorf("Expected updated CPU value to be '55%%', got %q", panel.Items[0].Value)
	}

	// Test updating non-existent item (should not panic)
	panel.UpdateItem("NonExistent", "value")
}

func TestInfoPanelClear(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	items := []InfoItem{
		{Label: "CPU", Value: "45%"},
		{Label: "Memory", Value: "2.5GB"},
	}
	panel.SetItems(items)

	panel.Clear()

	if len(panel.Items) != 0 {
		t.Errorf("Expected empty items after clear, got %d items", len(panel.Items))
	}
}
