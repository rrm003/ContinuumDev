/*
   Slice :
       Slice uses array as underlying data structure.
*/

package main

import (
	"fmt"
)

func main() {

	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var slice1 = array[:]
	fmt.Println("slice1 :", slice1)
	var slice2 = array[0:]
	fmt.Println("slice2 :", slice2)
	var slice3 = array[3:5]
	fmt.Println("slice3 :", slice3)
	var slice4 = array[6:len(array)]
	fmt.Println("slice4 :", slice4)
	
	//Invalid 
	//var slice5 = array[-6:-4]
	//fmt.Println("slice5 :", slice5)
	
}
