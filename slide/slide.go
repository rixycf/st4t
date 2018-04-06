package slide

import (

	// "image"
	"fmt"

	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	yaml "gopkg.in/yaml.v2"
)

const (
	fontfile = "./fonts/RictyDiminished-Regular.ttf"
)

var (
	dpi           = 72.0
	titleFontSize = 120.0
	bodyFontSize  = 80.0
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
	imgWidth, imgHeight int               //size of slide
	// Slide
	content Slide
	// size
}

func (s *SlideWriter) Init(w, h int) error {
	// init image size
	s.imgWidth, s.imgHeight = w, h

	s.img = image.NewRGBA(image.Rect(0, 0, s.imgWidth, s.imgHeight))
	draw.Draw(s.img, s.img.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)

	// init c freetypa.Context struct
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	s.c = freetype.NewContext()
	s.c.SetDPI(dpi)
	s.c.SetFont(f)
	s.c.SetFontSize(titleFontSize)
	s.c.SetClip(s.img.Bounds())
	s.c.SetDst(s.img)
	s.c.SetSrc(image.NewUniform(color.Black))

	// printf debug

	return nil
}

// read contents from yaml file
func (s *SlideWriter) ReadContents(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &s.content)
	if err != nil {
		return err
	}

	return nil
}

func (s *SlideWriter) DrawTitle(title string) error {
	pt := freetype.Pt(100, 100)
	s.c.SetFontSize(120)
	_, err := s.c.DrawString(title, pt)
	if err != nil {
		return err
	}
	return nil

}

func (s *SlideWriter) DrawBody(body string, color string) error {
	pt := freetype.Pt(200, 200)
	// s.c.SetSrc(image.NewUniform())
	if src, ok := colorMap[color]; ok {
		s.c.SetSrc(image.NewUniform(src))
	}
	_, err := s.c.DrawString(body, pt)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlideWriter) DrawImage(path string) error {
	// image file open
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return err
	}

	srcImg, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	pt := image.Point{X: -s.imgWidth / 2, Y: -s.imgHeight / 2}
	fmt.Printf("%+v\n", pt)
	// pt := image.Point{X: 20, Y: 20}
	fmt.Println(srcImg.Bounds())
	// draw src image on SlideWrite image
	draw.Draw(s.img, s.img.Bounds(), srcImg, pt, draw.Src)

	return nil
}

// return image data
func (s *SlideWriter) Render(wr io.Writer, w, h int) error {
	err := s.Init(w, h)
	if err != nil {
		return err
	}

	err = s.DrawTitle("concurrency in go ")
	if err != nil {
		return err
	}
	err = s.DrawImage("./image/Captain-falcon.png")
	if err != nil {
		return err
	}

	return png.Encode(wr, s.img)

}
