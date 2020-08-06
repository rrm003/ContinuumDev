package main

type pair struct {
	x int
	y int
}

func array1(p [2]pair) {
	p[0].x = -1
}
func array2(p [2]pair) {
	for i := range p {
		p[i].x = -1
	}
}
func array3(p [2]pair) {
	for _, pa := range p {
		pa.x = -1
	}
}
func array4(p []pair) {
	for _, pa := range p {
		pa.x = -1
	}
}
func array5(p []pair) {
	for i := range p {
		p[i].x = -1
	}
}

func main() {
	var arr1 [2]pair
	arr1[0] = pair{x: 1, y: 2}
	arr1[1] = pair{x: 3, y: 4}
	arr2 := []pair{pair{x: 1, y: 2},
		pair{x: 3, y: 4},
	}
	resetValue := func() {
		arr1[0] = pair{x: 1, y: 2}
		arr1[1] = pair{x: 3, y: 4}
		arr2 = []pair{pair{x: 1, y: 2},
			pair{x: 3, y: 4},
		}
	}
	// array1(arr1)
	// resetValue() //assume this function restores original value of arr1
	// array2(arr1)
	// resetValue() //assume this function restores original value of arr1
	// array3(arr1)
	// resetValue() //assume this function restores original value of arr1
	// array4(arr1) //will this compile ?
	// resetValue() //assume this function restores original value of arr1
	// array5(arr1) //will this compile ?
	array1(arr2) //will this compile ?
	resetValue() //assume this function restores original value of arr2
	array2(arr2) //will this compile ?
	resetValue() //assume this function restores original value of arr2
	array3(arr2) //will this compile ?
	resetValue() //assume this function restores original value of arr2
	array4(arr2) //will this compile ?
	resetValue() //assume this function restores original value of arr2
	array5(arr2) //will this compile ?
}
