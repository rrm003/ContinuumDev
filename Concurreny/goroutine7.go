package main

import (
	"fmt"
	"time"
)

func P(i int) {
	fmt.Println(i)
}

func O(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go O(i)
		go P(i)
	}
	time.Sleep(time.Millisecond)
}
