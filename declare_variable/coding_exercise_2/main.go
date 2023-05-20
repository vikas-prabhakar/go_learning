/*
Change the code from the previous exercise in the following way:

1. Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.

2. Declare x, y and z on a single line -> multiple short declarations

3. Remove the statement that prints out the variables. See the error!

4. Change the program to run without error using the blank identifier (_)
*/
package main

func main() {
	var (a int
		 b float64
		 c bool
		 d string)
	x, y, z := 15, 15.5, "Gopher!"
	_, _, _, _, _, _, _ = a, b, c, d, x, y, z
}