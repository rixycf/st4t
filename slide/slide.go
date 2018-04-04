package slide

import (
	// "fmt"
	// "image"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"

	"github.com/golang/freetype"
	yaml "gopkg.in/yaml.v2"
)

const (
	fontfile = "../fonts/RictyDiminished-Regular.ttf"
)

var (
	dpi           = 72.0
	titleFontSize = 120.0
	bodyFontSize  = 80.0
	// wonb          = true
	// spacing  = 5.0
)

type Slide struct {
	Title string
	Body  []struct {
		Color string
		Text  string
	}
	Image string
}

type SlideWriter struct {
	img                 *image.RGBA       // slide Image
	c                   *freetype.Context // setting of fontwriter
	imgWidth, imgHeight int
	// size
}

func NewSlideWriter() *SlideWriter {
	s := SlideWriter{}
	// init image
	s.img = image.NewRGBA(image.Rect(0, 0, s.imgWidth, s.imgHeight))
	draw.Draw(s.img, s.img.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)

	// init c freetypa.Context struct
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
	}

	s.c = freetype.NewContext()
	s.c.SetDPI(dpi)
	s.c.SetFont(f)
	s.c.SetFontSize(titleFontSize)
	s.c.SetClip(s.img.Bounds())
	s.c.SetDst(s.img)
	return &s

}
func (s *SlideWriter) init() {

}

func (s *SlideWriter) drawTitle(title string) error {
	pt := freetype.Pt(10, 10)
	s.c.SetFontSize(120)
	_, err := s.c.DrawString(title, pt)
	if err != nil {
		return err
	}
	return nil

}

func (s *SlideWriter) drawBody(body string) {

}

func (s *SlideWriter) drawImage(image string) {
}

// func NewSlide() *Slide {
//
// 	return &s
// }

// func (s *Slide) ymlUnmarshal(path string) error {
// 	err := yaml.Unmarshal(path, &s)
// 	return err
// }

// return image data
func SlideDrawer(path string) {
	s := Slide{}
	fmt.Println(s)

	filedata, err := ioutil.ReadFile(path)
	if err != nil {
	}

	err = yaml.Unmarshal(filedata, &s)
	if err != nil {
	}
	fmt.Println(s)

}
