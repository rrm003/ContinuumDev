package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello! This is golang world")

	b := make([]byte, 8)
	str := ""
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("n = %v err = %v b = %q \n", n, err, b)
		str = str + string(b[:n])
	}
	fmt.Println(str)
}
