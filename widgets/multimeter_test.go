package widgets

import (
	"testing"

	"github.com/deadjoe/termdodo/theme"
	"github.com/gdamore/tcell/v2"
)

func TestNewMultiMeter(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	x, y := 5, 10
	width, height := 20, 15
	mm := NewMultiMeter(screen, x, y, width, height)

	// Check initial values
	if mm.X != x || mm.Y != y {
		t.Errorf("Expected position (%d,%d), got (%d,%d)", x, y, mm.X, mm.Y)
	}
	if mm.Width != width || mm.Height != height {
		t.Errorf("Expected size (%d,%d), got (%d,%d)", width, height, mm.Width, mm.Height)
	}
	if mm.Screen != screen {
		t.Error("Screen not properly set")
	}
	if mm.Items == nil {
		t.Error("Items slice should be initialized")
	}
	if mm.Orientation != Horizontal {
		t.Error("Default orientation should be Horizontal")
	}
}

func TestMultiMeterItems(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	mm := NewMultiMeter(screen, 0, 0, 20, 10)

	// Test SetItems
	items := []MeterItem{
		{Label: "CPU", Value: 50, MaxValue: 100},
		{Label: "Memory", Value: 75, MaxValue: 100},
	}
	mm.SetItems(items)
	if len(mm.Items) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(mm.Items))
	}

	// Test AddItem
	newItem := MeterItem{Label: "Disk", Value: 30, MaxValue: 100}
	mm.AddItem(newItem)
	if len(mm.Items) != len(items)+1 {
		t.Error("AddItem failed to add new item")
	}
	if mm.Items[len(mm.Items)-1].Label != newItem.Label {
		t.Error("AddItem failed to add correct item")
	}

	// Test ClearItems
	mm.ClearItems()
	if len(mm.Items) != 0 {
		t.Error("ClearItems failed to clear items")
	}
}

func TestMultiMeterDisplay(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	mm := NewMultiMeter(screen, 0, 0, 20, 10)

	// Test show/hide controls
	mm.SetShowLabels(false)
	if mm.ShowLabels {
		t.Error("SetShowLabels failed")
	}

	mm.SetShowValues(false)
	if mm.ShowValues {
		t.Error("SetShowValues failed")
	}

	mm.SetShowBorder(true)
	if !mm.ShowBorder {
		t.Error("SetShowBorder failed")
	}
}

func TestMultiMeterLayout(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	mm := NewMultiMeter(screen, 0, 0, 20, 10)

	// Test orientation
	mm.SetOrientation(Vertical)
	if mm.Orientation != Vertical {
		t.Error("SetOrientation failed")
	}

	// Test dimensions
	labelWidth := 10
	mm.SetLabelWidth(labelWidth)
	if mm.LabelWidth != labelWidth {
		t.Error("SetLabelWidth failed")
	}

	meterHeight := 2
	mm.SetMeterHeight(meterHeight)
	if mm.MeterHeight != meterHeight {
		t.Error("SetMeterHeight failed")
	}

	spacing := 1
	mm.SetSpacing(spacing)
	if mm.Spacing != spacing {
		t.Error("SetSpacing failed")
	}
}

func TestMultiMeterStyle(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	mm := NewMultiMeter(screen, 0, 0, 20, 10)

	// Test widget style
	style := tcell.StyleDefault.Background(tcell.ColorRed)
	mm.SetStyle(style)
	if mm.Style != style {
		t.Error("SetStyle failed")
	}

	// Test label style
	labelStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue)
	mm.SetLabelStyle(labelStyle)
	if mm.LabelStyle != labelStyle {
		t.Error("SetLabelStyle failed")
	}
}

func TestMultiMeterUpdate(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	mm := NewMultiMeter(screen, 0, 0, 20, 10)
	item := MeterItem{Label: "CPU", Value: 50, MaxValue: 100}
	mm.AddItem(item)

	// Test updating existing meter
	newValue := 75.0
	mm.UpdateMeter("CPU", newValue)
	if mm.Items[0].Value != newValue {
		t.Error("UpdateMeter failed to update value")
	}

	// Test updating non-existent meter
	mm.UpdateMeter("NonExistent", 100)
	if len(mm.Items) != 1 {
		t.Error("UpdateMeter should not add new items")
	}
}

func TestMultiMeterDimensions(t *testing.T) {
	t.Parallel()
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	width, height := 20, 10
	mm := NewMultiMeter(screen, 0, 0, width, height)

	// Test GetWidth
	if got := mm.GetWidth(); got != width {
		t.Errorf("GetWidth() = %v, want %v", got, width)
	}

	// Test GetHeight
	if got := mm.GetHeight(); got != height {
		t.Errorf("GetHeight() = %v, want %v", got, height)
	}
}

func TestMultiMeterSetOrientation(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	screen.Init()
	defer screen.Fini()

	mm := NewMultiMeter(screen, 0, 0, 40, 10)

	// Test horizontal orientation
	mm.SetOrientation(Horizontal)
	if mm.Orientation != Horizontal {
		t.Errorf("Expected orientation to be Horizontal, got %v", mm.Orientation)
	}

	// Test vertical orientation
	mm.SetOrientation(Vertical)
	if mm.Orientation != Vertical {
		t.Errorf("Expected orientation to be Vertical, got %v", mm.Orientation)
	}
}

func TestMultiMeterSetLabelWidth(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	screen.Init()
	defer screen.Fini()

	mm := NewMultiMeter(screen, 0, 0, 40, 10)

	// Test setting label width
	mm.SetLabelWidth(15)
	if mm.LabelWidth != 15 {
		t.Errorf("Expected label width to be 15, got %v", mm.LabelWidth)
	}

	// Test setting zero label width
	mm.SetLabelWidth(0)
	if mm.LabelWidth != 0 {
		t.Errorf("Expected label width to be 0, got %v", mm.LabelWidth)
	}
}

func TestMultiMeterSetMeterHeight(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	screen.Init()
	defer screen.Fini()

	mm := NewMultiMeter(screen, 0, 0, 40, 10)

	// Test setting meter height
	mm.SetMeterHeight(2)
	if mm.MeterHeight != 2 {
		t.Errorf("Expected meter height to be 2, got %v", mm.MeterHeight)
	}

	// Test setting zero meter height (should default to 1)
	mm.SetMeterHeight(0)
	if mm.MeterHeight != 1 {
		t.Errorf("Expected meter height to be 1, got %v", mm.MeterHeight)
	}
}

func TestMultiMeterSetSpacing(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	screen.Init()
	defer screen.Fini()

	mm := NewMultiMeter(screen, 0, 0, 40, 10)

	// Test setting spacing
	mm.SetSpacing(2)
	if mm.Spacing != 2 {
		t.Errorf("Expected spacing to be 2, got %v", mm.Spacing)
	}

	// Test setting zero spacing
	mm.SetSpacing(0)
	if mm.Spacing != 0 {
		t.Errorf("Expected spacing to be 0, got %v", mm.Spacing)
	}
}

func TestMultiMeterBoundaryConditions(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	screen.Init()
	defer screen.Fini()

	mm := NewMultiMeter(screen, 0, 0, 40, 10)

	// Test adding item with empty label
	item := MeterItem{
		Label:    "",
		Value:    50,
		MaxValue: 100,
		Style:    theme.GetStyle(theme.ColorToHex(theme.Current.MainFg), theme.ColorToHex(theme.Current.MainBg)),
		Height:   1,
	}
	mm.AddItem(item)
	if len(mm.Items) != 0 {
		t.Error("Expected item with empty label to not be added")
	}

	// Test adding valid item
	validItem := MeterItem{
		Label:    "Test",
		Value:    50,
		MaxValue: 100,
		Style:    theme.GetStyle(theme.ColorToHex(theme.Current.MainFg), theme.ColorToHex(theme.Current.MainBg)),
		Height:   1,
	}
	mm.AddItem(validItem)
	if len(mm.Items) != 1 {
		t.Error("Expected valid item to be added")
	}

	// Test setting negative dimensions
	mm.SetMeterHeight(-1)
	if mm.MeterHeight != 1 {
		t.Error("Expected meter height to be 1")
	}

	mm.SetLabelWidth(-1)
	if mm.LabelWidth != 0 {
		t.Error("Expected label width to be 0")
	}

	mm.SetSpacing(-1)
	if mm.Spacing != 0 {
		t.Error("Expected spacing to be 0")
	}
}
