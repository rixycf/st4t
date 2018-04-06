package main

import (
	"fmt"
	"os"

	"github.com/rixycf/gsft/slide"
	"github.com/rixycf/gsft/term"
)

func main() {
	fmt.Println("vim-go")
	render()
}

func render() {
	// get treminal size
	size, err := term.Size()
	fmt.Println(size)
	if err != nil {
		fmt.Printf("term.Size function error :%v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("%+v", size)
	// width, height := size.Width, size.Height
	s := slide.SlideWriter{}
	// s.Init(size.Width, size.Height)

	term := &term.ImageWriter{}
	defer term.Close()

	if err := s.Render(term, size.Width, size.Height); err != nil {
		fmt.Printf("can't render image : %v\n", err)
		os.Exit(1)
	}
}
