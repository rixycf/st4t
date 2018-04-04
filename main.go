package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rixycf/gsft/slide"
	"github.com/rixycf/gsft/term"
)

func main() {
	fmt.Println("vim-go")
	s := slide.NewSlideWriter()
	fmt.Printf("%+v", s)
	// render()
	// slide.SlideDrawer("./yml_slides/1.yml")
}

func render() {
	// get treminal size
	size, err := term.Size()
	if err != nil {
		fmt.Printf("term.Size function error :%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v", size)
	// width, height := size.Width, size.Height

	term := &term.ImageWriter{}
	defer term.Close()
	// 画像を作る
	filedata, _ := ioutil.ReadFile("./Captain-falcon.png")

	term.Write(filedata)
}
