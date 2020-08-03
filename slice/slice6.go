//slice as parameter 

package main 

import("fmt")

func square(a []int){
    for i := range a {
        a[i] = a[i]*a[i]
    }
}

func main() {
    var slice []int = []int{1,2,3,4,5}
    fmt.Println(slice)
    square(slice)
    fmt.Println(slice)
}
