package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `_\|/` {
			fmt.Printf("\r%c", r)
			//time.Sleep(delay / 100)
		}
	}
}

func fibnumber(x int) int {
	if x < 2 {
		return x
	}
	return fibnumber(x-1) + fibnumber(x-2)
}

func main() {
	fmt.Println(fibnumber(45))
	go spinner(time.Millisecond)
	fmt.Println("..exiting main()")
}
