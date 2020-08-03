package main

import (
	"fmt"
	"os"
)

func main() {
	//create file
	//If file exists then it will overwrite
	f1, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error Creating file")
	}
	fmt.Println(f1, "File created!")
}
