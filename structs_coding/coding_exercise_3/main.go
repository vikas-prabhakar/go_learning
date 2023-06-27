/*
Iterate and print out the favorite colors of the struct value called me.
*/
package main

import "fmt"

type  person struct {
	name string
	age int
	colors []string	

}
func main(){
	
	me := person{name:"vikas",age:33,colors:[]string{"Red","Green"}}
	you := person{name:"Udemy",age:23,colors:[]string{"Black","White"}}

	_ = you
	for i, val := range me.colors {
		_ = i
		fmt.Println(val)
	}
}