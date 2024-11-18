package widgets

import (
    "testing"
    "github.com/gdamore/tcell/v2"
)

func TestNewGraph(t *testing.T) {
    // Create a new screen simulation
    screen := tcell.NewSimulationScreen("")
    if err := screen.Init(); err != nil {
        t.Fatal(err)
    }

    // Test graph creation
    g := NewGraph(screen, 1, 1, 10, 5)
    if g == nil {
        t.Fatal("Expected non-nil graph")
    }

    // Check initial values
    if g.X != 1 || g.Y != 1 {
        t.Errorf("Expected position (1,1), got (%d,%d)", g.X, g.Y)
    }
    if g.Width != 10 || g.Height != 5 {
        t.Errorf("Expected size (10,5), got (%d,%d)", g.Width, g.Height)
    }
    if len(g.Data) != 0 {
        t.Errorf("Expected empty data, got %d items", len(g.Data))
    }
}

func TestGraphSetData(t *testing.T) {
    screen := tcell.NewSimulationScreen("")
    if err := screen.Init(); err != nil {
        t.Fatal(err)
    }

    g := NewGraph(screen, 1, 1, 10, 5)
    
    // Test setting data
    testData := []float64{1.0, 2.0, 3.0}
    g.SetData(testData)
    
    if len(g.Data) != len(testData) {
        t.Errorf("Expected %d data points, got %d", len(testData), len(g.Data))
    }
    
    for i, v := range testData {
        if g.Data[i] != v {
            t.Errorf("Data point %d: expected %f, got %f", i, v, g.Data[i])
        }
    }
}

func TestGraphStyles(t *testing.T) {
    screen := tcell.NewSimulationScreen("")
    if err := screen.Init(); err != nil {
        t.Fatal(err)
    }

    g := NewGraph(screen, 1, 1, 10, 5)
    
    // Test each graph style
    styles := []GraphStyle{
        GraphStyleBraille,
        GraphStyleBlock,
        GraphStyleTTY,
    }
    
    for _, style := range styles {
        g.SetGraphStyle(style)
        if g.GraphStyle != style {
            t.Errorf("Expected graph style %v, got %v", style, g.GraphStyle)
        }
        
        // Set some test data and ensure Draw() doesn't panic
        g.SetData([]float64{1.0, 2.0, 3.0})
        func() {
            defer func() {
                if r := recover(); r != nil {
                    t.Errorf("Draw() panicked with style %v: %v", style, r)
                }
            }()
            g.Draw()
        }()
    }
}

func TestGraphClear(t *testing.T) {
    screen := tcell.NewSimulationScreen("")
    if err := screen.Init(); err != nil {
        t.Fatal(err)
    }

    g := NewGraph(screen, 1, 1, 10, 5)
    g.SetData([]float64{1.0, 2.0, 3.0})
    
    g.Clear()
    if len(g.Data) != 0 {
        t.Errorf("Expected empty data after Clear(), got %d items", len(g.Data))
    }
}
