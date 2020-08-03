//len() cap()

package main 

import "fmt" 

func main() {

    var month [6]string = [6]string{"Jan", "Feb", "Mar", "Apr", "May", "June"} 
    var slice1 []string = month[2:5:6]
    fmt.Println(slice1,len(slice1),cap(slice1))
    var slice2 []string = month[:]
    fmt.Println(slice2,len(slice2),cap(slice2))
    var slice3 []string = append(slice2,"July", "Aug")
    fmt.Println(slice3,len(slice3),cap(slice3))
}
