# Termdodo Examples

This directory contains examples demonstrating the usage of various Termdodo widgets and features.

## Running Examples

Each example can be run directly using the Go command:

```bash
go run examples/basic/main.go
```

## Available Examples

### Basic Example
`basic/main.go` - Demonstrates basic usage of box and graph widgets.
- Box creation and styling
- Graph widget with different styles
- Theme integration

### Gradient Example
`gradient/main.go` - Shows gradient color effects in widgets.
- Gradient color generation
- Color transitions
- Multiple color schemes

### InfoPanel Example
`infopanel/main.go` - Demonstrates the InfoPanel widget.
- Key-value information display
- Dynamic updates
- Flexible layout

### MultiMeter Example
`multimeter/main.go` - Shows multiple meters in a single widget.
- Multiple meter management
- Individual labels and values
- Synchronized updates

### StatusBar Example
`statusbar/main.go` - Demonstrates the status bar widget.
- Multiple status items
- Dynamic updates
- Fixed position layout

### Table Example
`table/main.go` - Shows the table widget functionality.
- Column headers
- Sortable columns
- Scrollable content
- Row selection

### TreeView Example
`treeview/main.go` - Demonstrates the tree view widget.
- Hierarchical data display
- Node expansion/collapse
- Navigation
- Selection handling

## Key Controls

Most examples support the following keyboard controls:
- `q` or `Esc` - Quit the example
- `↑`/`↓` - Navigate items (where applicable)
- `←`/`→` - Expand/collapse nodes (in TreeView)
- `Enter` - Select/activate items
- `PgUp`/`PgDn` - Page navigation (in Table)

## Notes

- Each example is self-contained and can be run independently
- Examples demonstrate best practices for using Termdodo
- All examples use the default theme by default
- Window size is handled automatically
