package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("1.PWD 2.Files")
	var opt int = 0
	fmt.Scan(&opt)
	switch opt {
	case 1:
		dir, err := exec.Command("pwd").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(dir))

	case 2:
		list, err := exec.Command("ls").Output()
		if err != nil {
			fmt.Println(err, "..exiting")
		} else {
			output := string(list[:])
			fmt.Println(output)
		}

	default:
		fmt.Println("Wrong options!")
	}
}
