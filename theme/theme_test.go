package theme

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewTheme(t *testing.T) {
	theme := NewTheme()
	if theme == nil {
		t.Fatal("Expected non-nil Theme")
	}

	// Check default values
	if theme.Background != tcell.ColorBlack {
		t.Errorf("Expected default background color to be Black, got %v", theme.Background)
	}
}

func TestSetBackground(t *testing.T) {
	theme := NewTheme()
	testColor := tcell.ColorBlue
	theme.SetBackground(testColor)

	if theme.Background != testColor {
		t.Errorf("Expected background color to be %v, got %v", testColor, theme.Background)
	}
}

func TestSetForeground(t *testing.T) {
	theme := NewTheme()
	testColor := tcell.ColorWhite
	theme.SetForeground(testColor)

	if theme.Foreground != testColor {
		t.Errorf("Expected foreground color to be %v, got %v", testColor, theme.Foreground)
	}
}

func TestSetAccent(t *testing.T) {
	theme := NewTheme()
	testColor := tcell.ColorGreen
	theme.SetAccent(testColor)

	if theme.Accent != testColor {
		t.Errorf("Expected accent color to be %v, got %v", testColor, theme.Accent)
	}
}

func TestSetBorder(t *testing.T) {
	theme := NewTheme()
	testColor := tcell.ColorYellow
	theme.SetBorder(testColor)

	if theme.Border != testColor {
		t.Errorf("Expected border color to be %v, got %v", testColor, theme.Border)
	}
}

func TestGetStyle(t *testing.T) {
	theme := NewTheme()
	theme.SetBackground(tcell.ColorBlack)
	theme.SetForeground(tcell.ColorWhite)

	style := theme.GetStyle()
	bg, fg, _ := style.Decompose()

	if bg != tcell.ColorBlack {
		t.Errorf("Expected style background to be Black, got %v", bg)
	}
	if fg != tcell.ColorWhite {
		t.Errorf("Expected style foreground to be White, got %v", fg)
	}
}

func TestGetAccentStyle(t *testing.T) {
	theme := NewTheme()
	theme.SetBackground(tcell.ColorBlack)
	theme.SetAccent(tcell.ColorGreen)

	style := theme.GetAccentStyle()
	bg, fg, _ := style.Decompose()

	if bg != tcell.ColorBlack {
		t.Errorf("Expected style background to be Black, got %v", bg)
	}
	if fg != tcell.ColorGreen {
		t.Errorf("Expected style foreground to be Green, got %v", fg)
	}
}

func TestGetBorderStyle(t *testing.T) {
	theme := NewTheme()
	theme.SetBackground(tcell.ColorBlack)
	theme.SetBorder(tcell.ColorYellow)

	style := theme.GetBorderStyle()
	bg, fg, _ := style.Decompose()

	if bg != tcell.ColorBlack {
		t.Errorf("Expected style background to be Black, got %v", bg)
	}
	if fg != tcell.ColorYellow {
		t.Errorf("Expected style foreground to be Yellow, got %v", fg)
	}
}

func TestDarkTheme(t *testing.T) {
	theme := DarkTheme()
	if theme == nil {
		t.Fatal("Expected non-nil DarkTheme")
	}

	if theme.Background != tcell.ColorBlack {
		t.Errorf("Expected dark theme background to be Black, got %v", theme.Background)
	}
}

func TestLightTheme(t *testing.T) {
	theme := LightTheme()
	if theme == nil {
		t.Fatal("Expected non-nil LightTheme")
	}

	if theme.Background != tcell.ColorWhite {
		t.Errorf("Expected light theme background to be White, got %v", theme.Background)
	}
}
