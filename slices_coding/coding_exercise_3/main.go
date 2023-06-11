/*
1. Declare a slice called nums with three float64 numbers.

2. Append the value 10.1 to the slice

3. In one statement append to the slice the values: 4.1,  5.5 and 6.6

4. Print out the slice

5. Declare a slice called n with two float64 values

6. Append n to nums

7. Print out the nums slice

*/
package main
 
import "fmt"
 
func main() {
    nums := []float64{1.2, 5.6, 4.6}
    nums = append(nums, 10.1)
    nums = append(nums,4.1,  5.5, 6.6)
    fmt.Printf("Nums slice %v\n", nums)
    n := []float64{1.1,2.2} 
    nums = append(nums,n...)
    fmt.Printf("Nums slice %v\n", nums)
 

 
}