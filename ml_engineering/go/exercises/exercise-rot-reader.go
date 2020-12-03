// https://tour.golang.org/methods/23
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	if err != nil {
		return n, err
	}
	input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	for i := range b {
		idx := bytes.Index(input, []byte{b[i]})
		if idx > -1 {
			b[i] = output[idx]
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
