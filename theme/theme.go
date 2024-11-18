package theme

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// Theme represents a complete color theme
type Theme struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Author      string      `json:"author"`
	Background  tcell.Color `json:"-"`
	Foreground  tcell.Color `json:"-"`
	MainBg      tcell.Color `json:"-"`
	MainFg      tcell.Color `json:"-"`
	Title       tcell.Color `json:"-"`
	Meter       []string    `json:"meter_bg"`
	Graph       []string    `json:"graph"`
	Border      tcell.Color `json:"-"`
	Selected    tcell.Color `json:"-"`
	HighlightBg tcell.Color `json:"-"`
	HighlightFg tcell.Color `json:"-"`
	Accent      tcell.Color `json:"-"`

	// JSON fields for serialization
	BackgroundHex  string `json:"background"`
	ForegroundHex  string `json:"foreground"`
	MainBgHex      string `json:"main_bg"`
	MainFgHex      string `json:"main_fg"`
	TitleHex       string `json:"title"`
	BorderHex      string `json:"border"`
	SelectedHex    string `json:"selected"`
	HighlightBgHex string `json:"highlight_bg"`
	HighlightFgHex string `json:"highlight_fg"`
	AccentHex      string `json:"accent"`
}

// SetBackground sets the background color
func (t *Theme) SetBackground(color tcell.Color) {
	t.Background = color
	t.BackgroundHex = ColorToHex(color)
}

// SetForeground sets the foreground color
func (t *Theme) SetForeground(color tcell.Color) {
	t.Foreground = color
	t.ForegroundHex = ColorToHex(color)
}

// SetAccent sets the accent color
func (t *Theme) SetAccent(color tcell.Color) {
	t.Accent = color
	t.AccentHex = ColorToHex(color)
}

// SetBorder sets the border color
func (t *Theme) SetBorder(color tcell.Color) {
	t.Border = color
	t.BorderHex = ColorToHex(color)
}

// GetStyle returns the default style for the theme
func (t *Theme) GetStyle() tcell.Style {
	style := tcell.StyleDefault
	if t.Background != 0 {
		style = style.Background(t.Background)
	}
	if t.Foreground != 0 {
		style = style.Foreground(t.Foreground)
	}
	return style
}

// GetAccentStyle returns the accent style for the theme
func (t *Theme) GetAccentStyle() tcell.Style {
	style := tcell.StyleDefault.Normal()
	if t.Accent != 0 {
		style = style.Foreground(t.Accent)
	}
	return style
}

// GetBorderStyle returns the border style for the theme
func (t *Theme) GetBorderStyle() tcell.Style {
	style := tcell.StyleDefault.Normal()
	if t.Border != 0 {
		style = style.Foreground(t.Border)
	}
	return style
}

// GetGradientStyle returns a style based on position in a gradient
func (t *Theme) GetGradientStyle(position float64) tcell.Style {
	style := tcell.StyleDefault
	if t.Background != 0 {
		style = style.Background(t.Background)
	}
	if t.Foreground != 0 {
		style = style.Foreground(t.Foreground)
	}
	return style
}

var (
	// ErrReadThemeFile is returned when the theme file cannot be read
	ErrReadThemeFile = errors.New("failed to read theme file")
	// ErrParseThemeFile is returned when the theme file cannot be parsed
	ErrParseThemeFile = errors.New("failed to parse theme file")
	// ErrMarshalTheme is returned when the theme cannot be marshaled
	ErrMarshalTheme = errors.New("failed to marshal theme")
	// ErrCreateThemeDir is returned when the theme directory cannot be created
	ErrCreateThemeDir = errors.New("failed to create theme directory")
	// ErrWriteThemeFile is returned when the theme file cannot be written
	ErrWriteThemeFile = errors.New("failed to write theme file")

	// Current holds the current theme
	Current Theme
)

// LoadDefaultTheme loads the default theme
func LoadDefaultTheme() {
	theme := NewTheme()
	theme.Background = tcell.ColorDefault
	theme.Foreground = tcell.ColorDefault
	theme.MainBg = tcell.ColorDefault
	theme.MainFg = tcell.ColorDefault
	theme.Title = tcell.ColorGreen
	theme.Border = tcell.ColorGray
	theme.Selected = tcell.ColorBlue
	theme.HighlightBg = tcell.ColorDarkGray
	theme.HighlightFg = tcell.ColorWhite
	theme.Accent = tcell.ColorBlue
	Current = *theme
}

// NewTheme creates a new theme with default values
func NewTheme() *Theme {
	theme := &Theme{
		Name:        "Default",
		Author:      "Termdodo",
		Description: "Default theme for Termdodo",
		Background:  tcell.ColorDefault,
		Foreground:  tcell.ColorDefault,
		MainBg:      tcell.ColorDefault,
		MainFg:      tcell.ColorDefault,
		Title:       tcell.ColorGreen,
		Border:      tcell.ColorYellow,
		Selected:    tcell.ColorBlue,
		HighlightBg: tcell.ColorDarkGray,
		HighlightFg: tcell.ColorWhite,
		Accent:      tcell.ColorGreen,
	}
	return theme
}

// LoadTheme loads a theme from a JSON file
func LoadTheme(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrReadThemeFile, err)
	}

	var theme Theme
	if err := json.Unmarshal(data, &theme); err != nil {
		return fmt.Errorf("%w: %v", ErrParseThemeFile, err)
	}

	// Convert hex colors to tcell.Color
	theme.Background = ParseHexColor(theme.BackgroundHex)
	theme.Foreground = ParseHexColor(theme.ForegroundHex)
	theme.MainBg = ParseHexColor(theme.MainBgHex)
	theme.MainFg = ParseHexColor(theme.MainFgHex)
	theme.Title = ParseHexColor(theme.TitleHex)
	theme.Border = ParseHexColor(theme.BorderHex)
	theme.Selected = ParseHexColor(theme.SelectedHex)
	theme.HighlightBg = ParseHexColor(theme.HighlightBgHex)
	theme.HighlightFg = ParseHexColor(theme.HighlightFgHex)
	theme.Accent = ParseHexColor(theme.AccentHex)

	Current = theme
	return nil
}

// SaveTheme saves the current theme to a JSON file
func SaveTheme(path string) error {
	// Convert tcell.Color to hex colors
	Current.BackgroundHex = ColorToHex(Current.Background)
	Current.ForegroundHex = ColorToHex(Current.Foreground)
	Current.MainBgHex = ColorToHex(Current.MainBg)
	Current.MainFgHex = ColorToHex(Current.MainFg)
	Current.TitleHex = ColorToHex(Current.Title)
	Current.BorderHex = ColorToHex(Current.Border)
	Current.SelectedHex = ColorToHex(Current.Selected)
	Current.HighlightBgHex = ColorToHex(Current.HighlightBg)
	Current.HighlightFgHex = ColorToHex(Current.HighlightFg)
	Current.AccentHex = ColorToHex(Current.Accent)

	data, err := json.MarshalIndent(Current, "", "    ")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMarshalTheme, err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateThemeDir, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("%w: %v", ErrWriteThemeFile, err)
	}

	return nil
}

// ParseHexColor parses a hex color string into tcell.Color
func ParseHexColor(hex string) tcell.Color {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return tcell.ColorDefault
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return tcell.ColorDefault
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return tcell.ColorDefault
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return tcell.ColorDefault
	}

	return RGBToColor(r, g, b)
}

// RGBToColor converts RGB values to tcell.Color
func RGBToColor(r, g, b uint64) tcell.Color {
	// Ensure values are within valid range
	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}
	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

// ColorToHex converts tcell.Color to hex string
func ColorToHex(color tcell.Color) string {
	r, g, b := color.RGB()
	// r, g, b are already in range [0, 255]
	return fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
}

// GetStyle returns a tcell.Style with the given foreground and background colors
func GetStyle(fg, bg string) tcell.Style {
	style := tcell.StyleDefault

	if fg != "" {
		style = style.Foreground(ParseHexColor(fg))
	}
	if bg != "" {
		style = style.Background(ParseHexColor(bg))
	}

	return style
}

// GetColor returns a tcell color from a color name
func GetColor(name string) tcell.Color {
	if strings.HasPrefix(name, "#") {
		return ParseHexColor(name)
	}

	// Handle named colors
	switch strings.ToLower(name) {
	case "black":
		return tcell.ColorBlack
	case "red":
		return tcell.ColorRed
	case "green":
		return tcell.ColorGreen
	case "yellow":
		return tcell.ColorYellow
	case "blue":
		return tcell.ColorBlue
	case "purple", "magenta":
		return tcell.ColorPurple
	case "teal", "cyan":
		return tcell.ColorTeal
	case "white":
		return tcell.ColorWhite
	default:
		return tcell.ColorDefault
	}
}

// GetGradientStyle returns a style based on position in a gradient
func GetGradientStyle(colors []string, position float64) tcell.Style {
	if len(colors) == 0 {
		return tcell.StyleDefault
	}

	if position <= 0 {
		return GetStyle(colors[0], "")
	}

	if position >= 1 {
		return GetStyle(colors[len(colors)-1], "")
	}

	// Calculate the segment and position within it
	segmentCount := float64(len(colors) - 1)
	segment := int(position * segmentCount)
	if segment >= len(colors)-1 {
		segment = len(colors) - 2
	}

	// Use the color at the calculated segment
	return GetStyle(colors[segment], "")
}
