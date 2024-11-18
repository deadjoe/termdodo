package widgets

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestNewTable(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	table := NewTable(screen, 0, 0, 80, 24)
	if table == nil {
		t.Fatal("Expected non-nil table")
	}

	if table.Width != 80 || table.Height != 24 {
		t.Errorf("Expected table dimensions to be 80x24, got %dx%d", table.Width, table.Height)
	}
}

func TestTableSetColumns(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	table := NewTable(screen, 0, 0, 80, 24)
	columns := []Column{
		{Title: "ID", Width: 10},
		{Title: "Name", Width: 20},
		{Title: "Value", Width: 15},
	}

	table.SetColumns(columns)

	if len(table.Columns) != len(columns) {
		t.Errorf("Expected %d columns, got %d", len(columns), len(table.Columns))
	}

	for i, col := range table.Columns {
		if col.Title != columns[i].Title {
			t.Errorf("Column %d: expected title %s, got %s", i, columns[i].Title, col.Title)
		}
		if col.Width != columns[i].Width {
			t.Errorf("Column %d: expected width %d, got %d", i, columns[i].Width, col.Width)
		}
	}
}

func TestTableSetRows(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}

	table := NewTable(screen, 0, 0, 80, 24)
	rows := [][]string{
		{"1", "Item 1", "100"},
		{"2", "Item 2", "200"},
	}

	table.SetRows(rows)

	if len(table.Rows) != len(rows) {
		t.Errorf("Expected %d rows, got %d", len(rows), len(table.Rows))
	}

	for i, row := range table.Rows {
		if len(row) != len(rows[i]) {
			t.Errorf("Row %d: expected %d cells, got %d", i, len(rows[i]), len(row))
		}
		for j, cell := range row {
			if cell != rows[i][j] {
				t.Errorf("Row %d, Cell %d: expected %s, got %s", i, j, rows[i][j], cell)
			}
		}
	}
}
