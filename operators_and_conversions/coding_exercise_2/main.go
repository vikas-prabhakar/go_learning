/*
Consider the following declarations: 

var i = 3
var f = 3.2
var s1, s2 = "3.14", "5"
Write a Go program that converts:

1. i to string (int to string)

2. s2 to int (string to int)

3. f to string (float64 to string)

4. s1 to float64  (string to float64)

5. Print the value and the type for each variable created after conversion.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"
	iToStr := strconv.Itoa(i)
	fmt.Printf("The value of the i is %q and type is %T\n",iToStr,iToStr)
	//

	sToint, err := strconv.Atoi(s2)
	_ = err
	fmt.Printf("The value of the s2 is %d and type is %T\n",sToint,sToint)
	fToStr := strconv.FormatFloat(f, 'f', 3, 64)
	//fToStr := fmt.Sprintf("%f",f) second method

	fmt.Printf("The value of the f is %q and type is %T\n",fToStr,fToStr)
	sToFloat, err := strconv.ParseFloat(s1,64) 
	_ = err
	
	fmt.Printf("The value of the s1 is %f and type is %T\n",sToFloat,sToFloat)

}