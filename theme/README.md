# Termdodo Theme System

The theme system in Termdodo allows you to customize the appearance of your terminal UI applications.

## Directory Structure

```
theme/
├── data/           # Theme JSON files
│   └── default.json
├── theme.go        # Theme management code
└── README.md       # This file
```

## Theme File Format

Themes are defined in JSON files with the following structure:

```json
{
    "name": "Theme Name",
    "description": "Theme description",
    "author": "Author name",
    "background": "#000000",
    "foreground": "#ffffff",
    "main_bg": "#000000",
    "main_fg": "#ffffff",
    "title": "#ff8f40",
    "meter_bg": [
        "#color1",
        "#color2",
        "#color3"
    ],
    "graph": [
        "#color1",
        "#color2",
        "#color3",
        "#color4"
    ],
    "border": "#border_color",
    "selected": "#selected_color",
    "highlight_bg": "#highlight_bg_color",
    "highlight_fg": "#highlight_fg_color"
}
```

## Usage

```go
import "github.com/deadjoe/termdodo/theme"

// Load the default theme
theme.LoadDefaultTheme()

// Load a custom theme
theme.LoadTheme("path/to/theme.json")

// Get a style for rendering
style := theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg)

// Get a gradient style (useful for graphs and meters)
gradientStyle := theme.GetGradientStyle(theme.Current.Graph, 0.5) // position 0.0-1.0
```

## Creating Custom Themes

1. Create a new JSON file in the `theme/data` directory
2. Follow the theme file format shown above
3. Use hex color codes for all colors (e.g., "#ff0000" for red)
4. Load your theme using `theme.LoadTheme()`

## Built-in Themes

- `default.json`: The default theme inspired by btop

## Contributing New Themes

1. Create your theme file in the `theme/data` directory
2. Ensure all required colors are defined
3. Test your theme with different widgets
4. Submit a pull request
