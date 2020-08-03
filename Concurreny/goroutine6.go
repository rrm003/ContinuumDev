package main

import (
	"fmt"
	"time"
)

func C(i int) {
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func D(i int) {
	for ; i < 105; i++ {
		fmt.Println(i)
	}
}

func main() {
	go C(0)
	go D(100)
	time.Sleep(time.Millisecond)
}
