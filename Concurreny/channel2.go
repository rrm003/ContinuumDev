package main

import "fmt"

func main() {
	//Buffered channel
	ch := make(chan int, 3)
	//fmt.Println(<-ch)
	ch <- 0
	fmt.Println(<-ch)
	ch <- 1
	ch <- 10
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 2
	ch <- 10
	ch <- 20
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// ch <- 3
	// fmt.Println(<-ch)

}
