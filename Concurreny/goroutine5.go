package main

import (
	"fmt"
	"time"
)

func P(i int) {
	fmt.Println("P", i)
	time.Sleep(time.Millisecond)
}

func Q(i int) {
	fmt.Println("Q", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go P(i)
		go Q(i)
	}
	time.Sleep(time.Millisecond)
}
