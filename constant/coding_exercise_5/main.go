/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
*/
package main
 
//import "math"
 
func main() {
	const a int = 7
	const b float64 = 5.6
	const c =  float64(a) * b
 
	x := 8
    _ = x
	//const xc int = x // variables belong to runtime
 
	//const noIPv6 = math.Pow(2, 128)  // functions calls belong to runtime
    
}