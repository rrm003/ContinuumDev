//slice declaration 


package main

import ("fmt")

func declaration1() {

    //slice doesn't need to specify the size 
    var slice1 []int
    slice1 = []int{1,2,3,4,5} 
    fmt.Println("Declaration1:\nSlice1 : ", slice1)
}

func declaration2() {
    
    //slice can be declared using array 
    var arr [5]int = [5]int{10,20,30,40,50}
    var slice2 []int = arr[:]
    fmt.Println("Declaration2:\nSlice2 : ", slice2)   
}

func declaration3() {

    //slice with capacity
    var arr [10]int = [10]int{55,66,77,88,99,11}
    var slice3 []int = arr[0:4:10]
    fmt.Println("Declaration3:\nSlice3 : ",slice3, cap(slice3))
    slice3 = append(slice3,1,2,3,4,5)
    fmt.Println("Declaration3:\nSlice3 : ",slice3, cap(slice3))
    slice3 = append(slice3,1,3)
    fmt.Println("Declaration3:\nSlice3 : ",slice3, cap(slice3))
    
    
/*
    NOTE : the capacity will be doubled as number of elements cross the capacity of the slice. Increase in capacity is twice the initial capacity 

*/
}

func declaration4() {

    //using make()
    //slice can be initialised at runtime using builtin function make()
    //it creates new slice and initialize vales with zero value of element type
    months := make([]string, 12)
    fmt.Println("Declaration4:\nmonths :",months)
    fmt.Println(len(months), cap(months))
}

func declaration5(){
    //sliceing a slice 
    var months []string = []string {
                                "Jan", "Feb", "Mar", "Apr", "May",
                                "June", "July", "Aug", "Sep",
                                 "Oct", "Nov", "Dec",
                            }
     fmt.Println("Declaration :\nSlice : ", months)
     var winter []string = months[8:]
     fmt.Println("Slice of slice : ", winter) 
}

func main() {

    declaration1()
    declaration2()
    declaration3()
    declaration4()
    declaration5()
}
