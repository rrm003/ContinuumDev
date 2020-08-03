package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Open() only acess already present file
	//if file is absent produces error
	f1, err := os.Open("temp1.txt")
	if err != nil {
		fmt.Println(err, "Error opening File !!")
	}
	fmt.Println(f1, "Opened File!!")
}
