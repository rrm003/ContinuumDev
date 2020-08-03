package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name string
	var age int
	var height float64

	f, err := os.Open("temp.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		_, err := fmt.Fscanf(f, "%s \t %d \t %f \n", &name, &age, &height)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Closing file")
			} else {
				fmt.Println(err, "\nError in file")
			}
			f.Close()
			os.Exit(1)
		}
		fmt.Printf("%s \t %d \t %f \n", name, age, height)
	}
}
