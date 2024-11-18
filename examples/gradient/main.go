package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"termdodo/draw"
	"termdodo/theme"
	"termdodo/widgets"
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
	mainBox.SetTitle("Gradient Block Progress Demo")
	mainBox.SetRound(true)

	// Create progress meters with different colors in the same hue
	meters := make([]*widgets.Meter, 3)
	for i := range meters {
		meters[i] = widgets.NewMeter(screen,
			mainBox.InnerX(),
			mainBox.InnerY()+i*2, // 使用正常的垂直间距
			mainBox.InnerWidth()-10)
		
		// Configure block style
		meters[i].SetBlockStyle(true)
		meters[i].SetBlockSpacing(1)
		meters[i].SetShowPercentage(true)
		
		// Set different gradients for each meter
		switch i {
		case 0:
			meters[i].SetGradient(
				tcell.NewRGBColor(0, 100, 255),  // Light blue
				tcell.NewRGBColor(0, 200, 255),  // Dark blue
			)
		case 1:
			meters[i].SetGradient(
				tcell.NewRGBColor(255, 100, 0),  // Light orange
				tcell.NewRGBColor(255, 200, 0),  // Dark orange
			)
		case 2:
			meters[i].SetGradient(
				tcell.NewRGBColor(0, 180, 0),    // Light green
				tcell.NewRGBColor(150, 255, 150), // Dark green
			)
		}
	}

	// Animation loop
	quit := make(chan struct{})
	go func() {
		progress := 0.0
		increasing := true
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if increasing {
					progress += 0.01
					if progress >= 1.0 {
						progress = 1.0
						increasing = false
					}
				} else {
					progress -= 0.01
					if progress <= 0.0 {
						progress = 0.0
						increasing = true
					}
				}

				// Update each meter with different patterns
				meters[0].SetValue(progress * 100)
				meters[1].SetValue((math.Sin(progress*math.Pi)*0.5 + 0.5) * 100)
				meters[2].SetValue(math.Pow(progress, 2) * 100)

				// Draw boxes and meters
				screen.Clear()
				mainBox.Draw()
				for _, m := range meters {
					m.Draw()
				}
				screen.Show()

			case <-quit:
				return
			}
		}
	}()

	// Handle input
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				close(quit)
				return
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}
