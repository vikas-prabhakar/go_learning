/*
There are an error in the following Go program. Try to identify the error, change the code and run the program without errors.

*/
package main
 
func main() {
    const x int = 10
 
    // declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8}	-> We cannot declare a slice constant.
    var m = []int{1: 3, 4: 5, 6: 8}
    _ = m
}