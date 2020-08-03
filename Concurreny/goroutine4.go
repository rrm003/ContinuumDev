package main

import (
	"fmt"
	"time"
)

func X() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//time.Sleep(time.Second)
}

func Y() {
	for i := 100; i < 105; i++ {
		fmt.Println(i)
	}
}

func main() {
	go X()
	go Y()
	time.Sleep(time.Millisecond)
}
