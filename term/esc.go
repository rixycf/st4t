package term

import "fmt"

const (
	//
	esc = "\033"

	Up                 Direction = "A"
	Down               Direction = "B"
	Forward            Direction = "C"
	Backward           Direction = "D"
	NextLine           Direction = "E"
	PreviousLine       Direction = "F"
	HorizontalAbsolute Direction = "G"
)

func HideCursor() {
	print(esc + "?25l")
}

func ShowCursor() {
	print(esc + "?25h")
}

func CursorMove(d Direction, n int) {
	fmt.Printf("%s%d%s", esc, n, d)
}

func CursorSavaPositon() {
	print(esc + "[s")
}

func CursorRestorePosition() {
	print(esc + "[u")
}
