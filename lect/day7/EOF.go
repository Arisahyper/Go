package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	e := errors.New("EOF")
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
		if err == e {
			fmt.Println("これではtrueにはならない")
			fmt.Println("アドレスで比較しているため")
		}
	}
}
