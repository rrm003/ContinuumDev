package main

import (
	"fmt"
	"time"
)

func readchannel(ch <-chan int) {
	fmt.Print(<-ch)
}

func writechannel(ch chan<- int) {
	ch <- 1
}

func main() {
	ch := make(chan int)
	go writechannel(ch)
	go readchannel(ch)
	time.Sleep(time.Millisecond)
}
