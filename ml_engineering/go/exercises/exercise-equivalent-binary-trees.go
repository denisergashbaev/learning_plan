// https://tour.golang.org/concurrency/8
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var w func(t *tree.Tree)
	w = func(t *tree.Tree) {
		if t == nil {
			return
		}
		w(t.Left)
		ch <- t.Value
		w(t.Right)
	}
	w(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch1_vals := make([]int, 0)
	for v := range ch1 {
		ch1_vals = append(ch1_vals, v)
	}

	ch2 := make(chan int)
	go Walk(t2, ch2)
	ch2_vals := make([]int, 0)
	for v := range ch2 {
		ch2_vals = append(ch2_vals, v)
	}

	return func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}(ch1_vals, ch2_vals)
}

func main() {
	fmt.Println("Same?:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same?:", Same(tree.New(1), tree.New(2)))
}
