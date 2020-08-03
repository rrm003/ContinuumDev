package main

import (
	"fmt"
	"time"
)

func main() {
	//Unbuffered channel
	ch := make(chan int)
	fmt.Println("length : ", len(ch))
	fmt.Println("capacity : ", cap(ch))
	go func() {
		ch <- 0
	}()
	fmt.Println("length : ", len(ch))
	fmt.Println("capacity : ", cap(ch))
	fmt.Println(<-ch)
	time.Sleep(time.Millisecond)
}
