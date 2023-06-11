/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main
 
import "fmt"
 
func main() {
    mySlice := []float64{1.2, 5.6}
 
    mySlice[2] = 6
 
    a := 10
    mySlice[0] = a
 
    mySlice[3] = 10.10
 
    mySlice = append(mySlice, a)
 
    fmt.Println(mySlice)
 
}
*/
package main
 
import "fmt"
 
func main() {
    mySlice := []float64{1.2, 5.6}
 
	mySlice[1] = 6
 
    a := 10.2
    mySlice[0] = a
 
   	mySlice[1] = 10.10
 
    mySlice = append(mySlice, a)
 
    fmt.Println(mySlice)
 
}

