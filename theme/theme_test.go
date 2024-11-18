package theme

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewTheme(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	if theme == nil {
		t.Fatal("Expected non-nil Theme")
	}

	// Check default values
	if theme.Background != tcell.ColorDefault {
		t.Errorf("Expected default background color to be Default, got %v", theme.Background)
	}
}

func TestSetBackground(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	testColor := tcell.ColorBlue
	theme.SetBackground(testColor)

	if theme.Background != testColor {
		t.Errorf("Expected background color to be %v, got %v", testColor, theme.Background)
	}
}

func TestSetForeground(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	testColor := tcell.ColorWhite
	theme.SetForeground(testColor)

	if theme.Foreground != testColor {
		t.Errorf("Expected foreground color to be %v, got %v", testColor, theme.Foreground)
	}
}

func TestSetAccent(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	testColor := tcell.ColorGreen
	theme.SetAccent(testColor)

	if theme.Accent != testColor {
		t.Errorf("Expected accent color to be %v, got %v", testColor, theme.Accent)
	}
}

func TestSetBorder(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	testColor := tcell.ColorYellow
	theme.SetBorder(testColor)

	if theme.Border != testColor {
		t.Errorf("Expected border color to be %v, got %v", testColor, theme.Border)
	}
}

func TestGetStyle(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	theme.SetBackground(tcell.ColorDefault)
	theme.SetForeground(tcell.ColorDefault)

	style := theme.GetStyle()
	fg, bg, _ := style.Decompose()

	if bg != tcell.ColorDefault {
		t.Errorf("Expected style background to be Default, got %v", bg)
	}
	if fg != tcell.ColorDefault {
		t.Errorf("Expected style foreground to be Default, got %v", fg)
	}
}

func TestGetAccentStyle(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	theme.SetAccent(tcell.ColorGreen)

	style := theme.GetAccentStyle()
	fg, _, _ := style.Decompose()

	if !fg.Valid() {
		t.Error("Expected foreground color to be valid")
	}
	if fg != tcell.ColorGreen {
		t.Errorf("Expected style foreground to be Green (%v), got %v", tcell.ColorGreen, fg)
	}
}

func TestGetBorderStyle(t *testing.T) {
	t.Parallel()
	theme := NewTheme()
	theme.SetBorder(tcell.ColorYellow)

	style := theme.GetBorderStyle()
	fg, _, _ := style.Decompose()

	if !fg.Valid() {
		t.Error("Expected foreground color to be valid")
	}
	if fg != tcell.ColorYellow {
		t.Errorf("Expected style foreground to be Yellow (%v), got %v", tcell.ColorYellow, fg)
	}
}
