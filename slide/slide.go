package slide

import (

	// "image"

	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	yaml "gopkg.in/yaml.v2"
)

const (
	// fontfile = "./fonts/RictyDiminished-Regular.ttf"
	fontfile = "./fonts/RictyDiminished-Bold.ttf"
)

var (
	// dpi           = 72.0
	dpi           = 72.0
	titleFontSize = 120.0
	bodyFontSize  = 40.0

	bg = colorMap["bg"]
	fg = colorMap["fg"]
)

type Slide struct {
	Title string
	Body  []struct {
		Color string
		Text  string
	}
	Image struct {
		Path string
		X    int
		Y    int
	}
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
	// s.imgWidth, s.imgHeight = 2880, 1800
	// s.imgWidth, s.imgHeight = 200, 200

	s.img = image.NewRGBA(image.Rect(0, 0, s.imgWidth, s.imgHeight))
	draw.Draw(s.img, s.img.Bounds(), image.NewUniform(bg), image.ZP, draw.Over)

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
	s.c.SetSrc(image.NewUniform(fg))

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

func (s *SlideWriter) DrawTitle() error {
	pt := freetype.Pt(100, 100)
	s.c.SetFontSize(titleFontSize)
	_, err := s.c.DrawString(s.content.Title, pt)
	if err != nil {
		return err
	}
	return nil

}

func (s *SlideWriter) DrawBody() error {
	pt := freetype.Pt(200, 200)
	s.c.SetFontSize(bodyFontSize)
	// s.c.SetSrc(image.NewUniform())
	for _, b := range s.content.Body {

		if src, ok := colorMap[b.Color]; ok {
			s.c.SetSrc(image.NewUniform(src))
		}
		_, err := s.c.DrawString(b.Text, pt)
		if err != nil {
			return err
		}

		pt.Y += s.c.PointToFixed(bodyFontSize)

	}
	return nil
}

func (s *SlideWriter) DrawImage(path string, x, y int) error {
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

	pt := image.Point{X: -x, Y: -y}

	draw.Draw(s.img, s.img.Bounds(), srcImg, pt, draw.Over)

	return nil
}

// Render write png image to wr
func (s *SlideWriter) Render(wr io.Writer, w, h int) error {
	fmt.Println(s)
	err := s.Init(w, h)
	if err != nil {
		return err
	}

	err = s.DrawTitle()
	if err != nil {
		return err
	}

	err = s.DrawBody()
	if err != nil {
		return err
	}

	if s.content.Image.Path != "" {
		fmt.Println("success")
		err = s.DrawImage(s.content.Image.Path, s.content.Image.X, s.content.Image.Y)
		if err != nil {
			return err
		}
	}

	return png.Encode(wr, s.img)

}
