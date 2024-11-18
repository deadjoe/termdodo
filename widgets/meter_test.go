package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewMeter(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	meter := NewMeter(screen, 0, 0, 50)
	if meter == nil {
		t.Fatal("Expected non-nil meter")
	}

	if meter.Width != 50 {
		t.Errorf("Expected meter width to be 50, got %d", meter.Width)
	}
}

func TestMeterSetValue(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	meter := NewMeter(screen, 0, 0, 50)

	testCases := []struct {
		value    float64
		expected float64
	}{
		{0.5, 0.5},
		{1.0, 1.0},
		{0.0, 0.0},
		{-0.1, 0.0}, // Should clamp to 0
		{1.1, 1.0},  // Should clamp to 1
	}

	for _, tc := range testCases {
		meter.SetValue(tc.value)
		if meter.Value != tc.expected {
			t.Errorf("SetValue(%f): expected %f, got %f", tc.value, tc.expected, meter.Value)
		}
	}
}

func TestMeterSetShowPercentage(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	meter := NewMeter(screen, 0, 0, 50)

	meter.SetShowPercentage(true)
	if !meter.ShowPct {
		t.Error("Expected ShowPercentage to be true")
	}

	meter.SetShowPercentage(false)
	if meter.ShowPct {
		t.Error("Expected ShowPercentage to be false")
	}
}

func TestMeterSetLabel(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	meter := NewMeter(screen, 0, 0, 50)

	testLabel := "CPU Usage"
	meter.SetLabel(testLabel)
	if meter.Label != testLabel {
		t.Errorf("Expected label to be %q, got %q", testLabel, meter.Label)
	}
}
