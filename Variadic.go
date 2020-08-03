package main

import ("fmt")

func greatest(num ...int) int {

	x := 0
	for _,i := range num {
        //fmt.Println(i)
		if x < i {
			x = i
		}
	}
	return x
}

func main() {

	fmt.Println(greatest(1, 2))
	fmt.Println(greatest(0))
	fmt.Println(greatest(1, 2, 3))
	fmt.Println(greatest(4, 3, 2, 1))

}