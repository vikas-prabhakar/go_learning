/*
Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}

Write a Go program that counts the number of positive even numbers in the array.
*/
package main

import "fmt"

func main(){
	var count = 0
	nums := [...]int{30, -1, -6, 90, -6}
	for _, value := range nums {

		if value%2 == 0 && value > 0{
			count ++
		}
	}
	fmt.Printf("number of positive even numbers: %d\n",count)
}