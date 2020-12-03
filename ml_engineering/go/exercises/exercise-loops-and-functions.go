//https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, new_z := 1.0, -1.0
	for math.Abs(new_z-z) > 0.00001 {
		z = new_z
		new_z = z - (z*z-x)/(2*z)
		//fmt.Println("intermediary", z, new_z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
