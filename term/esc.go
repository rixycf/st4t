package term

// import "fmt"

const (
	esc = "\033"
)

func HideCursor() {
	print(csi + "?25l")
}

func ShowCursor() {
	print(csi + "?25h")
}

func CursorSavaPositon() {
	print(esc + "[s")
}

func CursorRestorePosition() {
	print(esc + "[u")
}
