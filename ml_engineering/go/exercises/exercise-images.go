// https://tour.golang.org/methods/25
package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	//(uint8(y) + uint8(x)) / 2, uint8(x) * uint8(y), and x^y, uint8(x) ^ uint8(y)
	v := uint8(math.Pow(float64(x), float64(y)))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{width: 100, height: 200}
	pic.ShowImage(m)
}
