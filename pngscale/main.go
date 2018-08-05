package main

import (
	"errors"
	"flag"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

var (
	in     = flag.String("in", "", "Input PNG file")
	out    = flag.String("out", "", "Output PNG file")
	width  = flag.Int("width", 1000, "Width of output image")
	height = flag.Int("height", 500, "Height of output image")
)

func run() error {
	flag.Parse()
	if *in == "" {
		return errors.New("Missing input file")
	}
	if *out == "" {
		return errors.New("Missing output file")
	}
	inFile, err := os.Open(*in)
	if err != nil {
		return err
	}
	defer inFile.Close()
	outFile, err := os.Create(*out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	inImg, err := png.Decode(inFile)
	if err != nil {
		return err
	}

	outImg := image.NewRGBA(image.Rect(0, 0, *width, *height))

	inMin := inImg.Bounds().Min
	inMax := inImg.Bounds().Max
	outMin := outImg.Bounds().Min
	outMax := outImg.Bounds().Max
	for x := outMin.X; x < outMax.X; x++ {
		xScaled := float64(x-outMin.X) / float64(outMax.X-outMin.X) * float64(inMax.X-inMin.X)
		for y := outMin.Y; y < outMax.Y; y++ {
			yScaled := float64(y-outMin.Y) / float64(outMax.Y-outMin.Y) * float64(inMax.Y-inMin.Y)
			outImg.Set(x, y, inImg.At(int(math.Round(xScaled)), int(math.Round(yScaled))))
		}
	}

	if err := png.Encode(outFile, outImg); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
