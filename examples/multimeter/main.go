package main

import (
	"math/rand"
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

	// Create main box
	width, height := screen.Size()
	mainBox := draw.NewBox(screen, 1, 1, width-2, height/2-2)
	mainBox.SetTitle("Multi-Meter Demo")
	mainBox.SetRound(true)

	// Create multi-meter
	mm := widgets.NewMultiMeter(screen,
		mainBox.InnerX(),
		mainBox.InnerY(),
		mainBox.InnerWidth(),
		mainBox.InnerHeight())

	// Add meters with labels
	mm.AddItem(widgets.MeterItem{Label: "CPU", Value: 0.0})
	mm.AddItem(widgets.MeterItem{Label: "Memory", Value: 0.0})
	mm.AddItem(widgets.MeterItem{Label: "Disk", Value: 0.0})
	mm.AddItem(widgets.MeterItem{Label: "Network", Value: 0.0})

	// Configure display options
	mm.SetShowValues(true)
	mm.SetShowLabels(true)
	mm.SetSpacing(1)

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

	// Initialize random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Update loop
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			// Update meters with random values
			mm.UpdateMeter("CPU", r.Float64())
			mm.UpdateMeter("Memory", r.Float64())
			mm.UpdateMeter("Disk", r.Float64())
			mm.UpdateMeter("Network", r.Float64())

			// Draw everything
			screen.Clear()
			mainBox.Draw()
			mm.Draw()
			screen.Show()
		}
	}
}
