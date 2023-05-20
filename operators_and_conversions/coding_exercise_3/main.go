/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
*/
package main
 
import "fmt"
 
func main() {
    x, y := 4, 5.1
 
    z := x * int(y)
    fmt.Println(z)
 
    a := 5
    b := 6.2 * float64(a)
    fmt.Println(b)
 
}
