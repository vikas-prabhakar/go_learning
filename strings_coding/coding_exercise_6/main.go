/*
Consider the following string declaration:s := "你好 Go!"

1. Convert the string to a rune slice.

2. Print out the rune slice

3. Iterate over the rune slice and print out each index and rune in the rune slice
*/

package main
 
import (
	"fmt"

)
 
func main() {
	s := "你好 Go!"
	r := []rune(s)
	fmt.Printf("rune slice: %#v\n",r)
	for i, v := range r {
		fmt.Printf("index: %d and rune slice: %c\n",i,v)
	}

	
}