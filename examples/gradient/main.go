package main

import (
	"github.com/deadjoe/termdodo/draw"
	"github.com/deadjoe/termdodo/theme"
	"github.com/deadjoe/termdodo/widgets"
	"github.com/gdamore/tcell/v2"
	"math"
	"time"
)

// createMainBox creates and returns the main box widget
func createMainBox(screen tcell.Screen) *draw.Box {
	width, height := screen.Size()
	mainBox := draw.NewBox(screen, 1, 1, width-2, height/2-2)
	mainBox.SetTitle("Gradient Block Progress Demo")
	mainBox.SetRound(true)
	return mainBox
}

// createMeter creates a new meter with the specified gradient colors
func createMeter(screen tcell.Screen, mainBox *draw.Box, index int, startColor, endColor tcell.Color) *widgets.Meter {
	meter := widgets.NewMeter(screen,
		mainBox.InnerX(),
		mainBox.InnerY()+index*2,
		mainBox.InnerWidth()-10)

	meter.SetBlockStyle(true)
	meter.SetBlockSpacing(1)
	meter.SetShowPercentage(true)
	meter.SetGradient(startColor, endColor)

	return meter
}

// updateMeters updates the meter values with a sine wave animation
func updateMeters(meters []*widgets.Meter, t float64) {
	for i, meter := range meters {
		// Calculate phase shift for each meter
		phase := float64(i) * math.Pi / 4
		// Calculate value using sine wave
		value := (math.Sin(t+phase) + 1) * 50
		meter.SetValue(value)
	}
}

// drawUI draws all UI components
func drawUI(screen tcell.Screen, mainBox *draw.Box, meters []*widgets.Meter) {
	screen.Clear()
	mainBox.Draw()
	for _, meter := range meters {
		meter.Draw()
	}
	screen.Show()
}

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
	mainBox := createMainBox(screen)

	// Create progress meters with different colors
	meters := []*widgets.Meter{
		createMeter(screen, mainBox, 0, tcell.NewRGBColor(0, 100, 255), tcell.NewRGBColor(0, 200, 255)),
		createMeter(screen, mainBox, 1, tcell.NewRGBColor(255, 100, 0), tcell.NewRGBColor(255, 200, 0)),
		createMeter(screen, mainBox, 2, tcell.NewRGBColor(0, 180, 0), tcell.NewRGBColor(150, 255, 150)),
	}

	// Animation loop
	t := 0.0
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	// Run the application
	go func() {
		for range ticker.C {
			// Update meters
			updateMeters(meters, t)
			t += 0.1

			// Draw UI
			drawUI(screen, mainBox, meters)
		}
	}()

	// Wait for quit event
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}
