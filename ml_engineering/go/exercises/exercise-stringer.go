package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (v IPAddr) String() string {
	// https://stackoverflow.com/a/14426447/256002
	const ( // iota is reset to 0
		Sprintf = iota // c0 == 0
		Buffer  = iota // c1 == 1
	)
	switch Buffer {
	case Sprintf:
		return fmt.Sprintf("%v.%v.%v.%v", v[0], v[1], v[2], v[3])
	case Buffer:
		// yourbasic.org/golang/build-append-concatenate-strings-efficiently/
		var b strings.Builder
		//Grow(32)
		for _, p := range v {
			fmt.Fprintf(&b, "%v.", p)
		}
		s := b.String() // no copying
		s = s[:b.Len()-1]
		return s
	default:
		return ""
	}

}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
