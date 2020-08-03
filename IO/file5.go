package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f1, err := os.Open("tmp1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	readdata := make([]byte, 64)
	for {
		n, err := f1.Read(readdata)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", readdata[:n])
	}
}
