//https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for iy := range p {
		p[iy] = make([]uint8, dx)
		for ix := range p[iy] {
			p[iy][ix] = (uint8(iy) + uint8(ix)) / 2
		}
	}
	return p

}

func main() {
	pic.Show(Pic)
}
