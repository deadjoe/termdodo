package theme

import (
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Color represents a terminal color
type Color struct {
	Name  string `json:"name"`
	Color string `json:"color"` // Hex color code
}

// Theme represents a complete color theme
type Theme struct {
	Name         string            `json:"name"`
	Description  string           `json:"description"`
	Author       string           `json:"author"`
	Background   string           `json:"background"`
	Foreground   string           `json:"foreground"`
	MainBg       string           `json:"main_bg"`
	MainFg       string           `json:"main_fg"`
	Title        string           `json:"title"`
	Meter        []string         `json:"meter_bg"`
	Graph        []string         `json:"graph"`
	BorderColor  string           `json:"border"`
	Selected     string           `json:"selected"`
	HighlightBg  string           `json:"highlight_bg"`
	HighlightFg  string           `json:"highlight_fg"`
}

var (
	// Current holds the current theme
	Current Theme
)

// LoadDefaultTheme loads the default theme
func LoadDefaultTheme() {
	// Set default theme values
	Current = Theme{
		Name:        "Default",
		Description: "Default btop theme",
		Author:      "aristocratos",
		Background:  "#0a0e14",
		Foreground:  "#b3b1ad",
		MainBg:      "#0a0e14",
		MainFg:      "#b3b1ad",
		Title:       "#ff8f40",
		Meter:       []string{"#5ccc96", "#95e6cb", "#b4f9f8"},
		Graph:       []string{"#26a269", "#33d17a", "#2ec27e", "#5ccc96"},
		BorderColor: "#ffb454",
		Selected:    "#ffb454",
		HighlightBg: "#1c1f25",
		HighlightFg: "#ffffff",
	}

	// Get the current file's directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "Warning: Could not determine theme directory\n")
		return
	}

	// Try to load from file if available
	themePath := filepath.Join(filepath.Dir(filename), "data", "default.json")
	if _, err := os.Stat(themePath); err == nil {
		if err := LoadTheme(themePath); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not load theme file %s: %v\n", themePath, err)
			fmt.Fprintf(os.Stderr, "Using built-in default theme\n")
		}
	}
}

// LoadTheme loads a theme from a JSON file
func LoadTheme(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read theme file: %v", err)
	}

	var theme Theme
	if err := json.Unmarshal(data, &theme); err != nil {
		return fmt.Errorf("failed to parse theme file: %v", err)
	}

	Current = theme
	return nil
}

// SaveTheme saves the current theme to a JSON file
func SaveTheme(path string) error {
	data, err := json.MarshalIndent(Current, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal theme: %v", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create theme directory: %v", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write theme file: %v", err)
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

	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

// GetStyle returns a tcell.Style with the given foreground and background colors
func GetStyle(fg, bg string) tcell.Style {
	style := tcell.StyleDefault

	if fg != "" {
		fgColor := ParseHexColor(fg)
		style = style.Foreground(fgColor)
	}

	if bg != "" {
		bgColor := ParseHexColor(bg)
		style = style.Background(bgColor)
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

// NewTheme creates a new theme with default values
func NewTheme() *Theme {
	return &Theme{
		Name:        "Default",
		Author:      "Termdodo",
		Description: "Default theme for Termdodo",
		Background:  "#000000",
		Foreground:  "#FFFFFF",
		MainBg:      "#000000",
		MainFg:      "#FFFFFF",
		Title:       "#00FF00",
		Meter:       []string{"#5ccc96", "#95e6cb", "#b4f9f8"},
		Graph:       []string{"#26a269", "#33d17a", "#2ec27e", "#5ccc96"},
		BorderColor: "#666666",
		Selected:    "#0000FF",
		HighlightBg: "#1c1f25",
		HighlightFg: "#ffffff",
	}
}
