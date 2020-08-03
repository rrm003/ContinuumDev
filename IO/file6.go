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
	defer f1.Close()
	//io.Copy() writes content to os.Stdout from f1
	_, err1 := io.Copy(os.Stdout, f1)
	if err1 != nil {
		fmt.Println("Failed to copy..")
		os.Exit(1)
	}

	fmt.Println(f1.Name(), n)
}
