package main

import (
	"fmt"
	"math"
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

	// Create widgets
	width, height := screen.Size()

	// Create boxes for layout
	graphBox := draw.NewBox(screen, 1, 1, width/2-2, height/2-2)
	graphBox.SetTitle("Graph Demo")
	graphBox.SetRound(true)

	meterBox := draw.NewBox(screen, width/2+1, 1, width-2, height/2-2)
	meterBox.SetTitle("Meters Demo")
	meterBox.SetRound(true)

	// Create graph widget
	graph := widgets.NewGraph(screen, graphBox.InnerX(), graphBox.InnerY(),
		graphBox.InnerWidth(), graphBox.InnerHeight()-1)
	graph.SetGraphStyle(widgets.GraphStyleBlock)

	// Create meters
	meters := make([]*widgets.Meter, 3)
	for i := range meters {
		meters[i] = widgets.NewMeter(screen,
			meterBox.InnerX(),
			meterBox.InnerY()+i*2,
			meterBox.InnerWidth()-10)
	}

	// Animation loop
	quit := make(chan struct{})
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		t := 0.0

		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				// Update graph data
				data := make([]float64, graph.Width)
				for i := range data {
					x := t + float64(i)*0.2
					data[i] = 50 + 30*math.Sin(x)
				}
				graph.SetData(data)
				t += 0.2

				// Update meters with different patterns
				meters[0].SetValue(50 + 45*math.Sin(t*0.5))
				meters[1].SetValue(50 + 45*math.Sin(t*0.5+math.Pi/2))
				meters[2].SetValue(50 + 45*math.Sin(t*0.5+math.Pi))

				// Draw boxes
				screen.Clear()
				graphBox.Draw()
				meterBox.Draw()

				// Draw widgets
				graph.Draw()
				for _, meter := range meters {
					meter.Draw()
				}

				// Show screen
				screen.Show()
			}
		}
	}()

	// Event loop
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				close(quit)
				return
			}
		}
	}
}
