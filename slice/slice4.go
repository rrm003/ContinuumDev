//Using slice 

package main 

import "fmt" 

func main() {

    var slice []string = []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov", "Dec"} 
    
    //for _,month := range slice {
    //    fmt.Println(month)
    //}
    
    for i := range slice {
        fmt.Println(i+1,"\t",slice[i])
    }
    
}
