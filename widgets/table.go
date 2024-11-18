package widgets

import (
	"sort"
	"strings"

	"github.com/deadjoe/termdodo/theme"
	"github.com/gdamore/tcell/v2"
)

// Column represents a table column
type Column struct {
	Title     string
	Width     int
	MinWidth  int
	MaxWidth  int
	Alignment Alignment // Left, Right, Center
}

// Alignment represents text alignment
type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

// Table represents a table widget
type Table struct {
	X, Y          int
	Width, Height int
	Screen        tcell.Screen
	Style         tcell.Style
	HeaderStyle   tcell.Style
	SelectedStyle tcell.Style

	Columns     []Column
	Rows        [][]string
	SelectedRow int

	ShowHeader    bool
	ShowBorder    bool
	Sortable      bool
	SortColumn    int
	SortAscending bool
	HighlightRow  bool
	ScrollOffset  int
	VisibleRows   int
}

// NewTable creates a new table widget
func NewTable(screen tcell.Screen, x, y, width, height int) *Table {
	mainBg := theme.GetColor(theme.Current.MainBg)
	selectedColor := theme.GetColor(theme.Current.Selected)

	return &Table{
		X:             x,
		Y:             y,
		Width:         width,
		Height:        height,
		Screen:        screen,
		Style:         theme.GetStyle(theme.Current.MainFg, theme.Current.MainBg),
		HeaderStyle:   theme.GetStyle(theme.Current.Title, theme.Current.MainBg),
		SelectedStyle: tcell.StyleDefault.Foreground(selectedColor).Background(mainBg),
		ShowHeader:    true,
		ShowBorder:    true,
		Sortable:      true,
		HighlightRow:  true,
	}
}

// SetColumns sets the table columns
func (t *Table) SetColumns(columns []Column) {
	t.Columns = columns
	t.adjustColumnWidths()
}

// SetRows sets the table rows
func (t *Table) SetRows(rows [][]string) {
	t.Rows = rows
	if t.SortColumn >= 0 {
		t.sort()
	}
	t.adjustColumnWidths()
}

// AddRow adds a row to the table
func (t *Table) AddRow(row []string) {
	t.Rows = append(t.Rows, row)
	if t.SortColumn >= 0 {
		t.sort()
	}
	t.adjustColumnWidths()
}

// ClearRows clears all rows from the table
func (t *Table) ClearRows() {
	t.Rows = nil
	t.SelectedRow = 0
	t.ScrollOffset = 0
}

// SetSortColumn sets the column to sort by
func (t *Table) SetSortColumn(col int) {
	if !t.Sortable {
		return
	}
	if col < 0 || col >= len(t.Columns) {
		return
	}
	if t.SortColumn == col {
		t.SortAscending = !t.SortAscending
	} else {
		t.SortColumn = col
		t.SortAscending = true
	}
	t.sort()
}

// sort sorts the table rows by the current sort column
func (t *Table) sort() {
	if t.SortColumn < 0 || t.SortColumn >= len(t.Columns) {
		return
	}
	sort.Slice(t.Rows, func(i, j int) bool {
		if t.SortAscending {
			return t.Rows[i][t.SortColumn] < t.Rows[j][t.SortColumn]
		}
		return t.Rows[i][t.SortColumn] > t.Rows[j][t.SortColumn]
	})
}

// adjustColumnWidths adjusts column widths to fit the table width
func (t *Table) adjustColumnWidths() {
	if len(t.Columns) == 0 {
		return
	}

	// Calculate available width
	availableWidth := t.Width
	if t.ShowBorder {
		availableWidth -= 2 // Account for left and right borders
	}
	availableWidth -= len(t.Columns) - 1 // Account for column separators

	// First pass: set minimum widths and calculate total minimum width
	totalMinWidth := 0
	for i := range t.Columns {
		// Set minimum width based on column title and content
		minWidth := len(t.Columns[i].Title)
		for _, row := range t.Rows {
			if i < len(row) && len(row[i]) > minWidth {
				minWidth = len(row[i])
			}
		}
		if t.Columns[i].MinWidth > minWidth {
			minWidth = t.Columns[i].MinWidth
		}
		t.Columns[i].Width = minWidth
		totalMinWidth += minWidth
	}

	// If we have extra space, distribute it proportionally
	if extraWidth := availableWidth - totalMinWidth; extraWidth > 0 {
		// Calculate total content width for proportional distribution
		totalContentWidth := 0
		for i := range t.Columns {
			totalContentWidth += t.Columns[i].Width
		}

		// Distribute extra width proportionally
		remainingExtra := extraWidth
		for i := range t.Columns {
			if i == len(t.Columns)-1 {
				// Last column gets all remaining extra width
				t.Columns[i].Width += remainingExtra
			} else {
				// Calculate proportional extra width
				extra := (t.Columns[i].Width * extraWidth) / totalContentWidth
				t.Columns[i].Width += extra
				remainingExtra -= extra
			}

			// Apply maximum width constraint if set
			if t.Columns[i].MaxWidth > 0 && t.Columns[i].Width > t.Columns[i].MaxWidth {
				remainingExtra += t.Columns[i].Width - t.Columns[i].MaxWidth
				t.Columns[i].Width = t.Columns[i].MaxWidth
			}
		}
	}
}

// Draw draws the table on the screen
func (t *Table) Draw() {
	if len(t.Columns) == 0 {
		return
	}

	// Draw border if enabled
	startY := t.Y
	if t.ShowBorder {
		t.drawBorder()
		startY++
	}

	// Draw header if enabled
	if t.ShowHeader {
		t.drawHeader(startY)
		startY++
	}

	// Draw rows
	t.drawRows(startY)
}

// drawBorder draws the table border
func (t *Table) drawBorder() {
	style := t.Style

	// Draw top border
	t.Screen.SetContent(t.X, t.Y, '┌', nil, style)
	for x := t.X + 1; x < t.X+t.Width-1; x++ {
		t.Screen.SetContent(x, t.Y, '─', nil, style)
	}
	t.Screen.SetContent(t.X+t.Width-1, t.Y, '┐', nil, style)

	// Draw bottom border
	t.Screen.SetContent(t.X, t.Y+t.Height-1, '└', nil, style)
	for x := t.X + 1; x < t.X+t.Width-1; x++ {
		t.Screen.SetContent(x, t.Y+t.Height-1, '─', nil, style)
	}
	t.Screen.SetContent(t.X+t.Width-1, t.Y+t.Height-1, '┘', nil, style)

	// Draw side borders
	for y := t.Y + 1; y < t.Y+t.Height-1; y++ {
		t.Screen.SetContent(t.X, y, '│', nil, style)
		t.Screen.SetContent(t.X+t.Width-1, y, '│', nil, style)
	}
}

// drawHeader draws the table header
func (t *Table) drawHeader(y int) {
	x := t.X
	if t.ShowBorder {
		x++
	}

	for i, col := range t.Columns {
		// Draw column title
		title := t.alignText(col.Title, col.Width, col.Alignment)
		style := t.HeaderStyle
		if t.Sortable && i == t.SortColumn {
			if t.SortAscending {
				title = title + string(0x25B2) // Unicode UP TRIANGLE
			} else {
				title = title + string(0x25BC) // Unicode DOWN TRIANGLE
			}
		}
		for i, r := range title {
			t.Screen.SetContent(x+i, y, r, nil, style)
		}

		// Draw column separator
		if i < len(t.Columns)-1 {
			t.Screen.SetContent(x+col.Width, y, '│', nil, t.Style)
			x += col.Width + 1
		}
	}
}

// drawRows draws the table rows
func (t *Table) drawRows(startY int) {
	visibleRows := t.Height - startY
	if t.ShowBorder {
		visibleRows--
	}

	maxRow := len(t.Rows)
	if t.ScrollOffset+visibleRows < maxRow {
		maxRow = t.ScrollOffset + visibleRows
	}

	for rowIdx := t.ScrollOffset; rowIdx < maxRow; rowIdx++ {
		y := startY + rowIdx - t.ScrollOffset
		x := t.X
		if t.ShowBorder {
			x++
		}

		row := t.Rows[rowIdx]
		style := t.Style
		if t.HighlightRow && rowIdx == t.SelectedRow {
			style = t.SelectedStyle
		}

		// Draw each cell in the row
		for i, col := range t.Columns {
			cellText := ""
			if i < len(row) {
				cellText = t.alignText(row[i], col.Width, col.Alignment)
			}

			// Draw cell content
			for i, r := range cellText {
				t.Screen.SetContent(x+i, y, r, nil, style)
			}
			// Fill remaining space in cell
			for i := len(cellText); i < col.Width; i++ {
				t.Screen.SetContent(x+i, y, ' ', nil, style)
			}

			// Draw column separator
			if i < len(t.Columns)-1 {
				t.Screen.SetContent(x+col.Width, y, '│', nil, t.Style)
				x += col.Width + 1
			}
		}
	}
}

// alignText aligns text within the given width
func (t *Table) alignText(text string, width int, alignment Alignment) string {
	if len(text) > width {
		return text[:width]
	}

	padding := width - len(text)
	switch alignment {
	case AlignLeft:
		return text + strings.Repeat(" ", padding)
	case AlignRight:
		return strings.Repeat(" ", padding) + text
	case AlignCenter:
		leftPad := padding / 2
		rightPad := padding - leftPad
		return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
	default:
		return text + strings.Repeat(" ", padding)
	}
}

// handleUpKey handles up arrow key event
func (t *Table) handleUpKey() bool {
	if t.SelectedRow > 0 {
		t.SelectedRow--
		// Scroll up if selected row is above visible area
		if t.SelectedRow < t.ScrollOffset {
			t.ScrollOffset = t.SelectedRow
		}
		return true
	}
	return false
}

// handleDownKey handles down arrow key event
func (t *Table) handleDownKey() bool {
	if t.SelectedRow < len(t.Rows)-1 {
		t.SelectedRow++
		// Scroll down if selected row is below visible area
		visibleRows := t.Height - 2 // Account for header and border
		if t.SelectedRow >= t.ScrollOffset+visibleRows {
			t.ScrollOffset = t.SelectedRow - visibleRows + 1
		}
		return true
	}
	return false
}

// handlePageUpKey handles page up key event
func (t *Table) handlePageUpKey() bool {
	visibleRows := t.Height - 2 // Account for header and border
	t.ScrollOffset -= visibleRows
	if t.ScrollOffset < 0 {
		t.ScrollOffset = 0
	}
	if t.SelectedRow > t.ScrollOffset+visibleRows-1 {
		t.SelectedRow = t.ScrollOffset + visibleRows - 1
	}
	return true
}

// handlePageDownKey handles page down key event
func (t *Table) handlePageDownKey() bool {
	visibleRows := t.Height - 2 // Account for header and border
	maxScroll := len(t.Rows) - visibleRows
	if maxScroll < 0 {
		maxScroll = 0
	}
	t.ScrollOffset += visibleRows
	if t.ScrollOffset > maxScroll {
		t.ScrollOffset = maxScroll
	}
	if t.SelectedRow < t.ScrollOffset {
		t.SelectedRow = t.ScrollOffset
	}
	return true
}

// handleHomeKey handles home key event
func (t *Table) handleHomeKey() bool {
	t.SelectedRow = 0
	t.ScrollOffset = 0
	return true
}

// handleEndKey handles end key event
func (t *Table) handleEndKey() bool {
	t.SelectedRow = len(t.Rows) - 1
	visibleRows := t.Height - 2
	t.ScrollOffset = len(t.Rows) - visibleRows
	if t.ScrollOffset < 0 {
		t.ScrollOffset = 0
	}
	return true
}

// HandleEvent handles keyboard events
func (t *Table) HandleEvent(ev *tcell.EventKey) bool {
	if len(t.Rows) == 0 {
		return false
	}

	switch ev.Key() {
	case tcell.KeyUp:
		return t.handleUpKey()
	case tcell.KeyDown:
		return t.handleDownKey()
	case tcell.KeyPgUp:
		return t.handlePageUpKey()
	case tcell.KeyPgDn:
		return t.handlePageDownKey()
	case tcell.KeyHome:
		return t.handleHomeKey()
	case tcell.KeyEnd:
		return t.handleEndKey()
	}

	return false
}

// GetSelectedRow returns the currently selected row
func (t *Table) GetSelectedRow() ([]string, int) {
	if t.SelectedRow >= 0 && t.SelectedRow < len(t.Rows) {
		return t.Rows[t.SelectedRow], t.SelectedRow
	}
	return nil, -1
}

// SetHighlightRow sets whether to highlight the selected row
func (t *Table) SetHighlightRow(highlight bool) {
	t.HighlightRow = highlight
}

// SetShowHeader sets whether to show the header
func (t *Table) SetShowHeader(show bool) {
	t.ShowHeader = show
}

// SetShowBorder sets whether to show the border
func (t *Table) SetShowBorder(show bool) {
	t.ShowBorder = show
}

// SetSortable sets whether the table is sortable
func (t *Table) SetSortable(sortable bool) {
	t.Sortable = sortable
}
