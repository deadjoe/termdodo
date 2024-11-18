package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/deadjoe/termdodo/draw"
	"github.com/deadjoe/termdodo/theme"
	"github.com/deadjoe/termdodo/widgets"
	"time"
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

	// Create main box
	width, height := screen.Size()
	mainBox := draw.NewBox(screen, 1, 1, width-2, height/2-2)
	mainBox.SetTitle("Info Panel Demo")
	mainBox.SetRound(true)

	// Create info panel
	panel := widgets.NewInfoPanel(screen,
		mainBox.InnerX(),
		mainBox.InnerY(),
		mainBox.InnerWidth(),
		mainBox.InnerHeight())

	// Add some info items
	panel.AddItem("CPU", "Intel i7-9700K")
	panel.AddItem("Memory", "32GB DDR4")
	panel.AddItem("Disk", "1TB NVMe SSD")
	panel.AddItem("OS", "Linux 5.15.0")
	panel.AddItem("Uptime", "2d 5h 30m")
	panel.AddItem("Load Avg", "1.25 0.75 0.50")

	// Main loop
	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quit)
					return
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	// Update loop
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			// Update some dynamic info
			panel.UpdateItem("Uptime", time.Now().Format("15:04:05"))
			
			// Draw everything
			screen.Clear()
			mainBox.Draw()
			panel.Draw()
			screen.Show()
		}
	}
}
