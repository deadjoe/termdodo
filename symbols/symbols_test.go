package symbols

import "testing"

func TestBoxDrawingSymbols(t *testing.T) {
	t.Parallel()

	// Test box drawing symbols
	tests := []struct {
		name     string
		got      rune
		expected rune
	}{
		{"BoxDrawingTopLeft", BoxDrawingTopLeft, '┌'},
		{"BoxDrawingTopRight", BoxDrawingTopRight, '┐'},
		{"BoxDrawingBottomLeft", BoxDrawingBottomLeft, '└'},
		{"BoxDrawingBottomRight", BoxDrawingBottomRight, '┘'},
		{"BoxDrawingHorizontal", BoxDrawingHorizontal, '─'},
		{"BoxDrawingVertical", BoxDrawingVertical, '│'},
		{"BoxDrawingRoundTopLeft", BoxDrawingRoundTopLeft, '╭'},
		{"BoxDrawingRoundTopRight", BoxDrawingRoundTopRight, '╮'},
		{"BoxDrawingRoundBottomLeft", BoxDrawingRoundBottomLeft, '╰'},
		{"BoxDrawingRoundBottomRight", BoxDrawingRoundBottomRight, '╯'},
		{"HLine", HLine, '─'},
		{"VLine", VLine, '│'},
		{"DottedVLine", DottedVLine, '╎'},
		{"TLCorner", TLCorner, '┌'},
		{"TRCorner", TRCorner, '┐'},
		{"BLCorner", BLCorner, '└'},
		{"BRCorner", BRCorner, '┘'},
		{"RoundLeftUp", RoundLeftUp, '╭'},
		{"RoundRightUp", RoundRightUp, '╮'},
		{"RoundLeftDown", RoundLeftDown, '╰'},
		{"RoundRightDown", RoundRightDown, '╯'},
		{"DivRight", DivRight, '┤'},
		{"DivLeft", DivLeft, '├'},
		{"DivUp", DivUp, '┬'},
		{"DivDown", DivDown, '┴'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %c, want %c", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestGraphSymbols(t *testing.T) {
	t.Parallel()

	// Test Braille patterns
	brailleTests := []struct {
		name     string
		got      rune
		expected rune
	}{
		{"BrailleStart", BrailleStart, '⠀'},
		{"Braille1", Braille1, '⡀'},
		{"Braille2", Braille2, '⣀'},
		{"Braille3", Braille3, '⣄'},
		{"Braille4", Braille4, '⣤'},
		{"Braille5", Braille5, '⣦'},
		{"Braille6", Braille6, '⣶'},
		{"Braille7", Braille7, '⣷'},
		{"BrailleFull", BrailleFull, '⣿'},
	}

	for _, tt := range brailleTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %c, want %c", tt.name, tt.got, tt.expected)
			}
		})
	}

	// Test Block patterns
	blockTests := []struct {
		name     string
		got      rune
		expected rune
	}{
		{"BlockStart", BlockStart, ' '},
		{"Block1", Block1, '▁'},
		{"Block2", Block2, '▂'},
		{"Block3", Block3, '▃'},
		{"Block4", Block4, '▄'},
		{"Block5", Block5, '▅'},
		{"Block6", Block6, '▆'},
		{"Block7", Block7, '▇'},
		{"BlockFull", BlockFull, '█'},
	}

	for _, tt := range blockTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %c, want %c", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestSymbolConsistency(t *testing.T) {
	t.Parallel()

	// Test that equivalent symbols have the same value
	if BoxDrawingTopLeft != TLCorner {
		t.Error("BoxDrawingTopLeft and TLCorner should be the same")
	}
	if BoxDrawingTopRight != TRCorner {
		t.Error("BoxDrawingTopRight and TRCorner should be the same")
	}
	if BoxDrawingBottomLeft != BLCorner {
		t.Error("BoxDrawingBottomLeft and BLCorner should be the same")
	}
	if BoxDrawingBottomRight != BRCorner {
		t.Error("BoxDrawingBottomRight and BRCorner should be the same")
	}
	if BoxDrawingHorizontal != HLine {
		t.Error("BoxDrawingHorizontal and HLine should be the same")
	}
	if BoxDrawingVertical != VLine {
		t.Error("BoxDrawingVertical and VLine should be the same")
	}
	if BoxDrawingRoundTopLeft != RoundLeftUp {
		t.Error("BoxDrawingRoundTopLeft and RoundLeftUp should be the same")
	}
	if BoxDrawingRoundTopRight != RoundRightUp {
		t.Error("BoxDrawingRoundTopRight and RoundRightUp should be the same")
	}
	if BoxDrawingRoundBottomLeft != RoundLeftDown {
		t.Error("BoxDrawingRoundBottomLeft and RoundLeftDown should be the same")
	}
	if BoxDrawingRoundBottomRight != RoundRightDown {
		t.Error("BoxDrawingRoundBottomRight and RoundRightDown should be the same")
	}
}
