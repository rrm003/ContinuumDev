package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `_\|/_` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibseries(x int) []int {
	i, j := 0, 1

	a := 0
	var f []int
	for i <= x {
		//temp := i
		f = append(f, i)
		i = j
		j = f[a] + j

		a = a + 1
	}
	return f
}

func main() {
	go spinner(time.Millisecond)
	fmt.Println(fibseries(1000000000000000))
}
