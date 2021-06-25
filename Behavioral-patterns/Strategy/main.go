package Strategy

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"os"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	Writer 	io.Writer
	LogWriter io.Writer
}

func(d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}
func(d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}


type ConsoleSquare struct{}

type ImageSquare struct {
	DestinationFilePath string
}

func (t *ImageSquare) SetLog(writer io.Writer) {
	panic("implement me")
}

func (t *ImageSquare) SetWriter(writer io.Writer) {
	panic("implement me")
}

func (c *ConsoleSquare) Print() error {
	println("Square")
	return nil
}

func (t *ImageSquare) Print() error {
	width := 800
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})
	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)
	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)
	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("Error opening image")
	}
	defer w.Close()
	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}
	return nil
}

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main(){
	flag.Parse()
	var activeStrategy PrintStrategy
	switch *output {
	case "console":
		activeStrategy = &TextSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &TextSquare{}
	}
	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}


type TextSquare struct {
	PrintOutput
}

/*func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("Circle"))
	io.Copy(t.Writer, r)
	return nil
}*/

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}

