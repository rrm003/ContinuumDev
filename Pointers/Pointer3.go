package main

import "fmt"

func reverse(arr *[5]int){
	length := len(arr)-1
	for i,j:=0,length; i<length; i,j = i+1,j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [5]int {10,20,30,40,50}
	fmt.Println(arr)
	reverse(&arr)
	fmt.Println(arr)
}