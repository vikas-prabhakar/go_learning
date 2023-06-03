/*
1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.

Print out the cities array and notice the value of its elements.

2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.

Print out the grades array and notice the value of its elements.

3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.

4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.

5. Iterate over grades using the range keyword and print out element by element.
*/
package main

import "fmt"

func main(){
	var cities [2]string
	fmt.Printf("%#v\n",cities) 
	var grades = [3]float64{1.23,2.12}
	fmt.Printf("%#v\n",grades) 
	var taskDone = [...]bool{false, true, false, true}
	fmt.Printf("%v\n",taskDone) 
	for i:=0 ; i < len(cities); i++ {
		fmt.Printf("%v\n",cities[i])
	}
	for i,value := range grades{
		_ = i
		fmt.Printf("%v\n",value)
	}
}