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
	// Walk over the tree and return the collected values
	f := func(t *tree.Tree) []int {
		ch := make(chan int)
		go Walk(t, ch)
		ch_vals := make([]int, 0)
		for v := range ch {
			ch_vals = append(ch_vals, v)
		}
		return ch_vals
	}

	// slice equality
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
	}(f(t1), f(t2))
}

func main() {
	fmt.Println("Same?:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same?:", Same(tree.New(1), tree.New(2)))
}
