package main

import (
	"fmt"
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
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	// Load default theme
	theme.LoadDefaultTheme()

	// Create main content box
	width, height := screen.Size()
	mainBox := draw.NewBox(screen, 1, 1, width-2, height-3)
	mainBox.SetTitle("Main Content")
	mainBox.SetRound(true)

	// Create status bar at the bottom
	statusBar := widgets.NewStatusBar(screen, 0, height-1, width)

	// Add some status items
	statusBar.AddItem(widgets.StatusItem{Text: "Ready", Style: tcell.StyleDefault})
	statusBar.AddItem(widgets.StatusItem{Text: "", Style: tcell.StyleDefault})
	statusBar.AddItem(widgets.StatusItem{Text: "Press 'q' to quit", Style: tcell.StyleDefault})

	// Main loop
	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
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
		}
	}()

	// Update loop
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			count++
			// Update status items
			statusBar.UpdateItem(1, time.Now().Format("15:04:05"))
			if count%5 == 0 {
				statusBar.UpdateItem(0, fmt.Sprintf("Processing... %d", count))
			}

			// Draw everything
			screen.Clear()
			mainBox.Draw()
			statusBar.Draw()
			screen.Show()
		}
	}
}
