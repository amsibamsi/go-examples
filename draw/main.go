package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

var (
	col = color.RGBA{255, 255, 255, 255}
	img = image.NewRGBA(image.Rect(0, 0, 500, 500))
)

func drawDot(x, y int) {
	img.Set(x, y, col)
	img.Set(x-1, y, col)
	img.Set(x, y-1, col)
	img.Set(x+1, y, col)
	img.Set(x, y+1, col)
}

func drawLine(x1, y1, x2, y2 float64) {
	dy := (y2 - y1) / (x2 - x1)
	dx := (x2 - x1) / (y2 - y1)
	if math.Abs(dx) >= math.Abs(dy) {
		y := y1
		for x := x1; x <= x2; x++ {
			img.Set(int(x), int(y), col)
			y += dy
		}
	} else {
		x := x1
		for y := y1; y <= y2; y++ {
			img.Set(int(x), int(y), col)
			x += dx
		}
	}
}

func main() {
	drawDot(1, 1)
	drawDot(499, 1)
	drawDot(499, 499)
	drawDot(1, 499)
	drawLine(1, 1, 499, 499)
	drawLine(1, 499, 499, 1)
	file, _ := os.Create("test.jpeg")
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
