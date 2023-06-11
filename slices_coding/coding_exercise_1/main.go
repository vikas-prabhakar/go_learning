/*
Using a composite literal declare and initialize a slice of type string with 3 elements.

Iterate over the slice and print each element in the slice and its index.
*/
package main


import "fmt"
func main(){
	city := []string{"Delhi", "Bengaluru", "Mumbai"}
	for i, v := range city {
		fmt.Printf("Index %d and city %q\n",i,v)
	}

}