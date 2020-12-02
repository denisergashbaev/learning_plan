// https://tour.golang.org/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := make(map[string]int)
	for _, v := range strings.Fields(s) {
		fields[v] += 1
	}
	return fields
}

func main() {
	wc.Test(WordCount)
}
