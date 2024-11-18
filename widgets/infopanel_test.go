package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewInfoPanel(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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

func TestInfoPanelSetFields(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	fields := []InfoField{
		{Label: "CPU", Value: "45%"},
		{Label: "Memory", Value: "2.5GB"},
		{Label: "Disk", Value: "120GB"},
	}
	panel.SetFields(fields)

	if len(panel.Fields) != len(fields) {
		t.Errorf("Expected %d fields, got %d", len(fields), len(panel.Fields))
	}

	for i, field := range fields {
		if panel.Fields[i].Label != field.Label || panel.Fields[i].Value != field.Value {
			t.Errorf("Field %d mismatch: expected %+v, got %+v", i, field, panel.Fields[i])
		}
	}
}

func TestInfoPanelAddField(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	label := "CPU"
	value := "45%"
	panel.AddField(label, value)

	if len(panel.Fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(panel.Fields))
	}

	field := panel.Fields[0]
	if field.Label != label || field.Value != value {
		t.Errorf("Field mismatch: expected {%q, %q}, got {%q, %q}", label, value, field.Label, field.Value)
	}
}

func TestInfoPanelClear(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	panel := NewInfoPanel(screen, 0, 0, 80, 10)
	panel.AddField("CPU", "45%")
	panel.AddField("Memory", "2.5GB")
	panel.ClearFields()

	if len(panel.Fields) != 0 {
		t.Errorf("Expected no fields after clear, got %d", len(panel.Fields))
	}
}
