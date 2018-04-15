package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rixycf/st4t/slide"
	"github.com/rixycf/st4t/term"
)

func main() {

	flag.Parse()
	dir := flag.Args()
	if len(dir) < 1 {
		fmt.Printf("please set args\n")
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

	// スライドショー return keyを押す毎にスライドを送る．
	slideShow(files)

}

func slideShow(path []string) {
	i := 0
	for {
		buf := make([]byte, 2)
		_, err := os.Stdin.Read(buf)
		if err == io.EOF {
			fmt.Printf("EOF")
			break
		}
		term.CursorSavaPositon()
		render(path[i])
		term.CursorRestorePosition()
		i++
		if i > len(path)-1 {
			i = 0
		}
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

func checkTerm() error {

	if os.Getenv("TERM") != "xterm-256color" {
		return errors.New("test")
	}

	if os.Getenv("TERM_PROGRAM") != "iTerm.app" {
		return errors.New("this app runs only on iTerm2")
	}
	return nil
}

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
