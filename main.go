package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rixycf/st4t/slide"
	"github.com/rixycf/st4t/term"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	flag.Parse()
	dir := flag.Args()
	if len(dir) < 1 {
		fmt.Printf("please set directory \n")
		os.Exit(1)
	}

	// check terminal. if terminal is not iTerm2, then exit program
	err := checkTerm()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	files, err := getYmlFiles(dir[0])
	fmt.Println(files)
	if err != nil {
		fmt.Printf("getYmlFiles() error : %v\n", err)
	}

	// スライドショー終了時にターミナルの履歴をクリアする
	defer term.ClearScrollback()

	err = cleanTerm()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	slideShow(files)

}

func cleanTerm() error {
	size, err := term.Size()
	if err != nil {
		return err
	}
	if size.Row == 0 {
		return errors.New("failed get terminal size")
	}
	print(strings.Repeat("\n", size.Row))
	term.CursorMove(term.Up, size.Row+2)
	return nil
}

// func (row int) error {
// 	if row == 0 {
// 		return errors.New("row value : 0")
// 	}
// 	print(strings.Repeat("\n", row))
// 	return nil
// }

// slideShow start slide show
// if you type return key , then this app show next slide
func slideShow(path []string) {
	// slide index
	i := 0
	const (
		hKey = 0x68
		jKey = 0x6a
		kKey = 0x6b
		lKey = 0x6c
		qKey = 0x71
	)
	term.CursorSavaPositon()
	render(path[i])
	term.CursorRestorePosition()

END_SLIDE:
	for {
		buf := make([]byte, 1)
		// 標準入力を一文字ずつ受け取るためにrawモードにする
		s, _ := terminal.MakeRaw(1)
		os.Stdin.Read(buf)
		terminal.Restore(1, s)

		term.CursorSavaPositon()

		// slide show
		switch buf[0] {
		case hKey:
			i--
			if i < 0 {
				i = len(path) - 1
			}
			render(path[i])
		case lKey:
			i++
			if i > len(path)-1 {
				i = 0
			}
			render(path[i])
		case qKey:
			break END_SLIDE
		default:
		}

		term.CursorRestorePosition()
	}
}

func render(path string) {
	// get treminal size
	size, err := term.Size()
	if err != nil {
		fmt.Printf("term.Size function error :%v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("%+v", size)
	// width, height := size.Width, size.Height
	s := slide.SlideWriter{}
	s.ReadContents(path)
	// s.Init(size.Width, size.Height)

	term := &term.ImageWriter{}
	defer term.Close()

	if err := s.Render(term, size.Width, size.Height); err != nil {
		fmt.Printf("can't render image : %v\n", err)
		os.Exit(1)
	}
}

// checkTerm check terminal application.
func checkTerm() error {

	if os.Getenv("TERM") != "xterm-256color" {
		return errors.New("this app runs only on iTerm2 ")
	}

	if os.Getenv("TERM_PROGRAM") != "iTerm.app" {
		return errors.New("this app runs only on iTerm2")
	}
	return nil
}

// getYmlFiles get yml files from args
func getYmlFiles(dir string) ([]string, error) {
	ymlFiles := make([]string, 0, 10)

	// get current directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dirname := filepath.Join(wd, dir)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		yp := filepath.Join(dirname, f.Name())
		ymlFiles = append(ymlFiles, yp)
	}

	return ymlFiles, nil
}
