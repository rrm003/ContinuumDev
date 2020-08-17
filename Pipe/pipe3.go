package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	packageReader, packageWriter := io.Pipe()

	go func() {
		defer packageWriter.Close()
		_, err := fmt.Fprintln(packageWriter, "hello")
		if err != nil {
			panic(err)
		}
	}()

	_, err := io.Copy(os.Stdout, packageReader)
	if err != nil {
		panic(err)
	}

}
