package main

import (
	"fmt"
	"os"
	"time"

	"github.com/deadjoe/termdodo/draw"
	"github.com/deadjoe/termdodo/theme"
	"github.com/deadjoe/termdodo/widgets"
	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}

	// Set up screen
	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	screen.Clear()

	// Handle cleanup
	defer func() {
		if r := recover(); r != nil {
			screen.Fini()
			fmt.Fprintf(os.Stderr, "Panic: %v\n", r)
			os.Exit(1)
		}
		screen.Fini()
		os.Exit(0)
	}()

	// Load default theme
	theme.LoadDefaultTheme()

	// Create main box
	width, height := screen.Size()
	mainBox := draw.NewBox(screen, 1, 1, width-2, height-2)
	mainBox.SetTitle("Table Demo")
	mainBox.SetRound(true)

	// Create table
	table := widgets.NewTable(screen,
		mainBox.InnerX(),
		mainBox.InnerY(),
		mainBox.InnerWidth(),
		mainBox.InnerHeight())

	// Set columns
	table.SetColumns([]widgets.Column{
		{Title: "ID", Width: 4, Alignment: widgets.AlignRight},
		{Title: "Name", Width: 20, Alignment: widgets.AlignLeft},
		{Title: "Age", Width: 5, Alignment: widgets.AlignRight},
		{Title: "City", Width: 15, Alignment: widgets.AlignLeft},
	})

	// Add some data
	data := [][]string{
		{"1", "John Doe", "30", "New York"},
		{"2", "Jane Smith", "25", "Los Angeles"},
		{"3", "Bob Johnson", "35", "Chicago"},
		{"4", "Alice Brown", "28", "Houston"},
		{"5", "Charlie Wilson", "32", "Phoenix"},
		{"6", "Diana Miller", "27", "Philadelphia"},
		{"7", "Edward Davis", "31", "San Antonio"},
		{"8", "Fiona Clark", "29", "San Diego"},
		{"9", "George White", "33", "Dallas"},
		{"10", "Helen Green", "26", "San Jose"},
	}
	table.SetRows(data)

	// Enable features
	table.SetSortable(true)
	table.SetHighlightRow(true)
	table.SetShowHeader(true)
	table.SetShowBorder(true)

	// Main loop
	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if table.HandleEvent(ev) {
					continue
				}
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					close(quit)
					return
				case tcell.KeyRune:
					if ev.Rune() == 'q' {
						close(quit)
						return
					}
				}
			case *tcell.EventResize:
				screen.Sync()
			}

			// Draw everything
			screen.Clear()
			mainBox.Draw()
			table.Draw()
			screen.Show()
		}
	}()

	<-quit
}
