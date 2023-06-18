/*
Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}

1. The above map declaration has an error. Solve the error!

2. Delete a key:value pair from the map.

3. Iterate over the map and print out each key and value.
*/

package main
import "fmt"
func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 3)
	for i, v := range m {
		fmt.Printf("Key of map m is: %d and value: %t \n",i,v) //%t	the word true or false
	}
}