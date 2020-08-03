package main

import (
	"fmt"
	"os"
)

func main() {
	//os.OpenFile() provides generic low-level functionality to open existing file
	//Take 3 arguments : path, access(r/w), posixcompliant permission value
	f1, err1 := os.OpenFile("tmp1.txt", os.O_RDONLY, 444)
	f2, err2 := os.OpenFile("tmp2.txt", os.O_WRONLY, 222) //doesnot create file
	f3, err3 := os.OpenFile("temp.txt", os.O_APPEND, 777)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(f1, "F1 Creted")
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(f2, "F2 created")
	}

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(f3, "F3 opened")
	}

}
