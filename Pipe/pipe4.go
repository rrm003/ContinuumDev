package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.StringWriter(os.Stdout, fmt.Scanner())
}
