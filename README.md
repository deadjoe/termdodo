# Termdodo - Terminal UI Library

A powerful, theme-based terminal UI library written in Go, inspired by btop system monitoring tool.

## Features

### Core Components

- **Box Drawing**
  - Rounded and square borders
  - Title support
  - Flexible styling
  - Theme integration

- **Widgets**
  - Graph Widget
    * Multiple styles (Braille, Block, TTY)
    * Dynamic data scaling
    * Gradient color support
    * Real-time updates
  - Meter Widget
    * Percentage-based visualization
    * Gradient color support
    * Configurable width
    * Dynamic updates
  - MultiMeter Widget
    * Multiple meters in one widget
    * Individual labels and values
    * Synchronized updates
  - InfoPanel Widget
    * Key-value information display
    * Dynamic updates
    * Flexible layout
  - StatusBar Widget
    * Multiple status items
    * Dynamic updates
    * Fixed position
  - Table Widget
    * Column headers
    * Sortable columns
    * Scrollable content
    * Row selection
  - TreeView Widget
    * Hierarchical data display
    * Expandable nodes
    * Node selection
    * Keyboard navigation

- **Theme System**
  - JSON-based configuration
  - Dynamic theme switching
  - Gradient color generation
  - Default theme included
  - Custom theme support

### Technical Features

- Cross-platform support
- Unicode-aware rendering
- Performant drawing primitives
- Modular architecture
- Event-driven updates
- Flexible layouts
- Comprehensive examples

## Installation

```bash
go get github.com/deadjoe/termdodo
```

## Dependencies

- github.com/gdamore/tcell/v2 v2.5.3
- github.com/mattn/go-runewidth v0.0.16

## Quick Start

```go
package main

import (
    "github.com/deadjoe/termdodo/widgets"
    "github.com/deadjoe/termdodo/theme"
    "github.com/deadjoe/termdodo/draw"
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

    // Create a box
    box := draw.NewBox(screen, 1, 1, 30, 10)
    box.SetTitle("My Box")
    box.SetRound(true)

    // Create a graph
    graph := widgets.NewGraph(screen, box.InnerX(), box.InnerY(), 
        box.InnerWidth(), box.InnerHeight())
    graph.SetGraphStyle(widgets.GraphStyleBlock)

    // Set data
    data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    graph.SetData(data)

    // Draw
    box.Draw()
    graph.Draw()
    screen.Show()
}
```

## Examples

Check out the `examples` directory for complete examples of each widget:

- `examples/basic` - Basic usage of box and graph widgets
- `examples/gradient` - Gradient color effects demo
- `examples/infopanel` - Information panel demo
- `examples/multimeter` - Multiple meter display
- `examples/statusbar` - Status bar usage
- `examples/table` - Table widget with sorting
- `examples/treeview` - Tree view navigation

## Widget Documentation

### Box
```go
box := draw.NewBox(screen, x, y, width, height)
box.SetTitle("Title")
box.SetRound(true)  // Use rounded corners
```

### Graph
```go
graph := widgets.NewGraph(screen, x, y, width, height)
graph.SetGraphStyle(widgets.GraphStyleBlock)
graph.SetData([]float64{1.0, 2.0, 3.0})
```

### Meter
```go
meter := widgets.NewMeter(screen, x, y, width)
meter.SetValue(0.75)  // 75%
meter.SetShowPercentage(true)
```

### MultiMeter
```go
mm := widgets.NewMultiMeter(screen, x, y, width, height)
mm.AddMeter("CPU", 0.5)
mm.AddMeter("Memory", 0.75)
```

### InfoPanel
```go
panel := widgets.NewInfoPanel(screen, x, y, width, height)
panel.AddItem("CPU", "Intel i7")
panel.AddItem("Memory", "32GB")
```

### StatusBar
```go
bar := widgets.NewStatusBar(screen, x, y, width)
bar.AddItem("status", "Ready", style)
bar.AddItem("time", "12:00", style)
```

### Table
```go
table := widgets.NewTable(screen, x, y, width, height)
table.SetHeaders([]string{"ID", "Name", "Value"})
table.SetData([][]string{{"1", "Item 1", "100"}})
```

### TreeView
```go
tree := widgets.NewTreeView(screen, x, y, width, height)
root := tree.AddRoot("Root")
child := tree.AddNode(root, "Child")
```

## Theme System

See the [theme documentation](theme/README.md) for details on customizing the appearance.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
