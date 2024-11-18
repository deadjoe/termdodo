package symbols

// Box drawing symbols
const (
	BoxDrawingTopLeft        = "┌"
	BoxDrawingTopRight       = "┐"
	BoxDrawingBottomLeft     = "└"
	BoxDrawingBottomRight    = "┘"
	BoxDrawingHorizontal     = "─"
	BoxDrawingVertical       = "│"
	BoxDrawingRoundTopLeft   = "╭"
	BoxDrawingRoundTopRight  = "╮"
	BoxDrawingRoundBottomLeft = "╰"
	BoxDrawingRoundBottomRight = "╯"
	HLine         = "─"
	VLine         = "│"
	DottedVLine   = "╎"
	LeftUp        = "┌"
	RightUp       = "┐"
	LeftDown      = "└"
	RightDown     = "┘"
	RoundLeftUp   = "╭"
	RoundRightUp  = "╮"
	RoundLeftDown = "╰"
	RoundRightDown= "╯"
	DivRight      = "┤"
	DivLeft       = "├"
	DivUp         = "┬"
	DivDown       = "┴"
)

// Graph symbols
const (
	// Braille patterns for graph drawing
	BrailleStart = '⠀' // Braille blank
	Braille1     = '⡀'
	Braille2     = '⣀'
	Braille3     = '⣄'
	Braille4     = '⣤'
	Braille5     = '⣦'
	Braille6     = '⣶'
	Braille7     = '⣷'
	BrailleFull  = '⣿'

	// Block patterns for graph drawing
	BlockStart = ' '      // Space
	Block1     = '▁'     // Lower one eighth block
	Block2     = '▂'     // Lower one quarter block
	Block3     = '▃'     // Lower three eighths block
	Block4     = '▄'     // Lower half block
	Block5     = '▅'     // Lower five eighths block
	Block6     = '▆'     // Lower three quarters block
	Block7     = '▇'     // Lower seven eighths block
	BlockFull  = '█'     // Full block

	// TTY patterns for graph drawing
	TTYStart = ' '
	TTY1     = '_'
	TTY2     = '.'
	TTY3     = '-'
	TTY4     = '='
	TTY5     = '+'
	TTY6     = '*'
	TTY7     = '#'
	TTYFull  = '@'
)

// Meter symbols
const (
	Meter = "█"
)

// Direction symbols
const (
	ArrowUp    = "↑"
	ArrowDown  = "↓"
	ArrowLeft  = "←"
	ArrowRight = "→"
	Up    = "↑"
	Down  = "↓"
	Left  = "←"
	Right = "→"
	Enter = "↵"
)

// BraillePatterns returns all braille patterns in order
var BraillePatterns = []string{
	string(BrailleStart),
	string(Braille1),
	string(Braille2),
	string(Braille3),
	string(Braille4),
	string(Braille5),
	string(Braille6),
	string(Braille7),
	string(BrailleFull),
}

// BlockPatterns returns all block patterns in order
var BlockPatterns = []string{
	string(BlockStart),
	string(Block1),
	string(Block2),
	string(Block3),
	string(Block4),
	string(Block5),
	string(Block6),
	string(Block7),
	string(BlockFull),
}

// TTYPatterns returns all TTY patterns in order
var TTYPatterns = []string{
	string(TTYStart),
	string(TTY1),
	string(TTY2),
	string(TTY3),
	string(TTY4),
	string(TTY5),
	string(TTY6),
	string(TTY7),
	string(TTYFull),
}

// SuperScript numbers
var SuperScript = []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}
