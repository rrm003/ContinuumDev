package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Pipe2..")
	scanner := bufio.NewScanner(os.Stdin)
	var variable string
	for scanner.Scan() {
		variable = scanner.Text()
		fmt.Println(variable)
	}

	info := syscall.Environ()
	var key string
	system := make(map[string]string)
	for i := range info {
		rslt := strings.Split(info[i], "=")
		if rslt[0] != "" || rslt[1] != "" {
			key = rslt[0]
			system[key] = rslt[1]
			fmt.Println(key)
		}
	}
	fmt.Println(system[variable])
}
