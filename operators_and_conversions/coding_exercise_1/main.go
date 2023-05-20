/*
Consider the following declarations: 

var i int = 3
var f float64 = 3.2
Write a Go program that converts i to float64 and f to int.

Print out the  v and the value of the variables created after conversion.
*/
package main

import "fmt"

func main(){
	var i int = 3
	var f float64 = 3.2
	var x = float64(i)
	var y = int(f)
	fmt.Printf("Type of the x is %T and value is %.3f\n",x,x)
	fmt.Printf("Type of the y is %T and value is %d\n",y,y)
}