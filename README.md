# Termdodo - Terminal UI Library

[![Go](https://github.com/deadjoe/termdodo/actions/workflows/go.yml/badge.svg)](https://github.com/deadjoe/termdodo/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/deadjoe/termdodo?v=1)](https://goreportcard.com/report/github.com/deadjoe/termdodo)
[![codecov](https://codecov.io/gh/deadjoe/termdodo/branch/main/graph/badge.svg)](https://codecov.io/gh/deadjoe/termdodo)
[![GoDoc](https://godoc.org/github.com/deadjoe/termdodo?status.svg)](https://godoc.org/github.com/deadjoe/termdodo)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/deadjoe/termdodo/blob/main/LICENSE)

A powerful, theme-based terminal UI library written in Go, inspired by btop system monitoring tool. Termdodo provides a rich set of widgets and drawing primitives for creating beautiful terminal user interfaces with minimal effort.

## Features

### Core Components

- **Drawing Primitives**
  - Box drawing with rounded and square borders
  - Primitive shapes and lines
  - Unicode symbol support
  - Gradient color effects
  - Theme integration

- **Widgets**
  - Graph Widget
    * Multiple styles (Block, Braille)
    * Dynamic data scaling
    * Gradient color support
    * Real-time updates
  - Meter Widget
    * Percentage-based visualization
    * Gradient color support
    * Configurable width and style
    * Dynamic updates
  - MultiMeter Widget
    * Multiple meters in one widget
    * Individual labels and values
    * Synchronized updates
    * Flexible layout
  - InfoPanel Widget
    * Key-value information display
    * Dynamic updates
    * Flexible layout
    * Custom styling
  - StatusBar Widget
    * Multiple status items
    * Dynamic updates
    * Fixed position
    * Theme support
  - Table Widget
    * Column headers
    * Sortable columns
    * Scrollable content
    * Row selection
    * Custom cell styling
  - TreeView Widget
    * Hierarchical data display
    * Expandable/collapsible nodes
    * Keyboard navigation
    * Node selection
    * Custom node styling

- **Theme System**
  - JSON-based configuration
  - Dynamic theme switching
  - Gradient color generation
  - Default theme included
  - Custom theme support
  - Style inheritance

### Technical Features

- Comprehensive test coverage
- Cross-platform support
- Unicode-aware rendering
- Performant drawing primitives
- Modular architecture
- Event-driven updates
- Flexible layouts
- Well-documented API
- Rich examples

## Installation

```bash
go get github.com/deadjoe/termdodo
```

## Quick Start

```go
package main

import (
    "github.com/deadjoe/termdodo/widgets"
    "github.com/deadjoe/termdodo/theme"
    "github.com/deadjoe/termdodo/draw"
    "github.com/gdamore/tcell/v2"
)

func main() {
    // Initialize screen
    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    if err := screen.Init(); err != nil {
        panic(err)
    }
    defer screen.Fini()

    // Load default theme
    theme.LoadDefaultTheme()

    // Create widgets
    box := draw.NewBox(screen, 1, 1, 30, 10)
    box.SetTitle("System Info")
    box.SetRound(true)

    info := widgets.NewInfoPanel(screen, box.InnerX(), box.InnerY(), 
        box.InnerWidth(), box.InnerHeight())
    info.AddItem("CPU", "Intel i7 2.6GHz")
    info.AddItem("Memory", "16GB / 32GB")
    info.AddItem("Disk", "256GB SSD")

    // Main loop
    for {
        box.Draw()
        info.Draw()
        screen.Show()

        // Handle events
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape {
                return
            }
        }
    }
}
```

## Examples

Check out the `examples` directory for complete examples of each widget:

- `examples/basic` - Basic usage of box and primitive drawings
- `examples/gradient` - Gradient color effects demonstration
- `examples/infopanel` - Information panel with dynamic updates
- `examples/multimeter` - Multiple meter display with real-time updates
- `examples/statusbar` - Status bar with multiple items
- `examples/table` - Table widget with sorting and selection
- `examples/treeview` - Tree view with keyboard navigation

## Package Documentation

### draw
The `draw` package provides primitive drawing functions and box drawing capabilities:
```go
// Create a new box
box := draw.NewBox(screen, x, y, width, height)
box.SetTitle("Title")
box.SetRound(true)

// Draw primitive shapes
draw.DrawLine(screen, x1, y1, x2, y2, style)
```

### widgets
The `widgets` package contains all UI widgets:
```go
// Create a graph widget
graph := widgets.NewGraph(screen, x, y, width, height)
graph.SetData([]float64{1.0, 2.0, 3.0})

// Create a meter widget
meter := widgets.NewMeter(screen, x, y, width)
meter.SetValue(0.75)

// Create a table widget
table := widgets.NewTable(screen, x, y, width, height)
table.SetHeaders([]string{"ID", "Name", "Value"})

// Create a tree view widget
tree := widgets.NewTreeView(screen, x, y, width, height)
root := tree.AddNode(nil, "Root")
```

### theme
The `theme` package handles theming and styling:
```go
// Load a custom theme
theme.LoadTheme("custom.json")

// Get style for a component
style := theme.GetStyle("widget.normal")
```

### symbols
The `symbols` package provides Unicode symbols for drawing:
```go
// Get box drawing symbols
boxChars := symbols.GetBoxCharacters(symbols.BoxStyleRound)
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## Testing

Run the test suite:
```bash
go test ./...
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
