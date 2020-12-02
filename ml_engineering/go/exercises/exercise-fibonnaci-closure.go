// https://tour.golang.org/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// hack so that iterations=0,1,2 return 0, 1, 1
	f0, f1 := -1, 1
	return func() int {
		f := f0 + f1
		f0, f1 = f1, f
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration %d: %d\n", i, f())
	}
}
