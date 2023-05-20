/*
Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.

Using short declaration syntax declare and assign these values to variables x, y and z:

- 20

- 15.5

- "Gopher!"

Using fmt.Println() print out the values of a, b, c, d, x, y and z.

Try to run the program without error.

*/
package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	x := 15
	y := 15.5
	z := "Gopher!"
	//fmt.Printf("a: %d \nb: %f \nc: %t\nd: %v\nx: %d\ny: %f\nz: %v", a,b,c,d,x,y,z)
	fmt.Println("a:",a,"\n","b: ",b,"\n","c: ",c,"\n","d: ",d,"\n","x: ",x,"\n","y: ",y,"\n","z: ",z)


}