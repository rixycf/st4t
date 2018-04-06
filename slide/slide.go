package slide

import (

	// "image"

	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
)

const (
// fontfile = "../fonts/RictyDiminished-Regular.ttf"
)

var (
	dpi           = 72.0
	titleFontSize = 120.0
	bodyFontSize  = 80.0
	// fontfile      = "./RictyDiminished-Regular.ttf"
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
	// Slide
	s Slide
	// size
}

// func NewSlideWriter() *SlideWriter {
// 	s := SlideWriter{}
// 	// init image
// 	s.img = image.NewRGBA(image.Rect(0, 0, s.imgWidth, s.imgHeight))
// 	draw.Draw(s.img, s.img.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)
//
// 	// init c freetypa.Context struct
// 	fontBytes, err := ioutil.ReadFile(fontfile)
// 	if err != nil {
// 	}
//
// 	f, err := freetype.ParseFont(fontBytes)
// 	if err != nil {
// 	}
//
// 	s.c = freetype.NewContext()
// 	s.c.SetDPI(dpi)
// 	s.c.SetFont(f)
// 	s.c.SetFontSize(titleFontSize)
// 	s.c.SetClip(s.img.Bounds())
// 	s.c.SetDst(s.img)
// 	return &s
//
// }

func (s *SlideWriter) Init(w, h int) error {
	// init image size
	s.imgWidth, s.imgHeight = w, h

	s.img = image.NewRGBA(image.Rect(0, 0, s.imgWidth, s.imgHeight))
	draw.Draw(s.img, s.img.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)

	// init c freetypa.Context struct
	fontBytes, err := ioutil.ReadFile("./fonts/RictyDiminished-Regular.ttf")
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

func (s *SlideWriter) DrawTitle(title string) error {
	pt := freetype.Pt(10, 10)
	s.c.SetFontSize(120)
	_, err := s.c.DrawString("aaa", pt)
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

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	pt := image.Point{X: s.imgWidth / 2, Y: s.imgHeight / 2}
	// draw src image on SlideWrite image
	draw.Draw(s.img, img.Bounds(), img, pt, draw.Over)

	return nil
}

// return image data
func (s *SlideWriter) Render(wr io.Writer, w, h int) error {
	err := s.Init(w, h)
	if err != nil {
		return err
	}
	err = s.DrawTitle("aaaaa")
	if err != nil {
		return err
	}

	return png.Encode(wr, s.img)

}
