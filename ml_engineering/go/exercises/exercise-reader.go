// https://tour.golang.org/methods/22
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
