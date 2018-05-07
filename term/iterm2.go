package term

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// escape sequence
var ecsi = "\033" // \033 -> \x1b -> ESC
var st = "\007"   // \007 -> \xa -> beep

var cellSizeOnce sync.Once
var cellWidth, cellHeight float64

type TermSize struct {
	Width  int
	Height int
	Col    int
	Row    int
}

func fileSetReadDeadline(f *os.File, t time.Time) error {
	return nil
}

func ClearScrollback() {
	print(ecsi + "]1337;ClearScrollback" + st)
}

func initCellSize() {
	// terminalをrawモードに変更する
	// 標準入力をそのままプロセスに渡すモード バッファリングしない
	s, err := terminal.MakeRaw(1)
	if err != nil {
		return
	}
	// fmt.Printf("value of : %+v\n", s)
	// fmt.Printf("type of : %T\n", s)
	defer terminal.Restore(1, s)
	// iTerm2のエスケープコードを使用してセルのサイズを取得
	fmt.Fprintf(os.Stdout, ecsi+"]1337;ReportCellSize"+st)
	// setread dead line set
	fileSetReadDeadline(os.Stdout, time.Now().Add(time.Second))
	defer fileSetReadDeadline(os.Stdout, time.Time{})
	fmt.Fscanf(os.Stdout, ecsi+"]1337;ReportCellSize=%f;%f"+ecsi+"\\", &cellHeight, &cellWidth)
}

func Size() (size TermSize, err error) {
	// size. 縦何列, 横何列かを取得
	size.Col, size.Row, err = terminal.GetSize(0)
	if err != nil {
		return
	}

	cellSizeOnce.Do(initCellSize)
	if cellWidth+cellHeight == 0 {
		err = errors.New("cannot get iterm2 cell size")
	}

	// ターミナルの幅は, 縦の列の数とそのセルの幅の積
	// ターミナルの高さは, 行の数とそのセルの高さとの積
	size.Width, size.Height = size.Col*int(cellWidth), size.Row*int(cellHeight)
	return
}
