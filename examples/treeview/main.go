package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/deadjoe/termdodo/draw"
	"github.com/deadjoe/termdodo/theme"
	"github.com/deadjoe/termdodo/widgets"
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
	mainBox := draw.NewBox(screen, 1, 1, width-2, height-2)
	mainBox.SetTitle("Tree View Demo")
	mainBox.SetRound(true)

	// Create tree view
	tree := widgets.NewTreeView(screen,
		mainBox.InnerX(),
		mainBox.InnerY(),
		mainBox.InnerWidth(),
		mainBox.InnerHeight())

	// Add some nodes
	root := tree.AddRoot("Project")
	
	// Add source code section
	src := tree.AddNode(root, "src")
	tree.AddNode(src, "main.go")
	tree.AddNode(src, "config.go")
	tree.AddNode(src, "utils.go")

	// Add tests section
	tests := tree.AddNode(root, "tests")
	tree.AddNode(tests, "main_test.go")
	tree.AddNode(tests, "config_test.go")
	tree.AddNode(tests, "utils_test.go")

	// Add docs section
	docs := tree.AddNode(root, "docs")
	tree.AddNode(docs, "README.md")
	tree.AddNode(docs, "API.md")
	tree.AddNode(docs, "CONTRIBUTING.md")

	// Add config files
	configs := tree.AddNode(root, "config")
	tree.AddNode(configs, "config.yaml")
	tree.AddNode(configs, "config.dev.yaml")
	tree.AddNode(configs, "config.prod.yaml")

	// Expand root node
	tree.ExpandNode(root)

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
				case tcell.KeyUp:
					tree.SelectPrevious()
				case tcell.KeyDown:
					tree.SelectNext()
				case tcell.KeyRight:
					tree.ExpandSelected()
				case tcell.KeyLeft:
					tree.CollapseSelected()
				case tcell.KeyEnter:
					tree.ToggleSelected()
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
			tree.Draw()
			screen.Show()
		}
	}()

	<-quit
}
