// https://tour.golang.org/methods/20
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// if we do not use float64(e), then the program runs into infinite
	// loop (executing it returns "Program exited: status 2.").
	// The reason is, imho, that fmt.Sprintf calls .Error() function
	// of the passed ErrorNegativeSqrt type which then recursively
	// calls to fmt.Sprintf() inside of it...
	return fmt.Sprintf("cannot Sqrt negative number: %s", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x)
	}
	z, new_z := 1.0, -1.0
	for math.Abs(new_z-z) > 0.00001 {
		z = new_z
		new_z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	for _, v := range []float64{2, -2} {
		if ret, err := Sqrt(v); err == nil {
			fmt.Println(ret)
		} else {
			fmt.Println(err)
		}
	}
}
