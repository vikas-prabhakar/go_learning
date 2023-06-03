/*

Using a for loop print out all the years from your birthday to the current year.

Use a variant of for loop where the post statement is moved inside the for block of code.
*/
package main

import (
    "fmt"
    "time"
)


func main(){
	birthDate := 1991
	curYear := time.Now().Year()
	for i := birthDate; i <= curYear; {
		fmt.Printf("%d \n",i)
		i++
	}
	
}