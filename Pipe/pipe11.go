package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var variable string
	fmt.Println("Pipe11..")
	fmt.Println("Enter Variable : ")
	fmt.Scanln(&variable)

	x := strings.NewReader(variable)
	io.Copy(os.Stdout, x)
}
