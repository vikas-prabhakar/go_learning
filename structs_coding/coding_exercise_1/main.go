/*
1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.

2. Declare and initialize two values of type person, one called me and another called you.

3. Print out the struct values.
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
	
	// %v	the value in a default format
	fmt.Printf("%v\n",me)
	// The %+v form shows the fields by name, while %#v formats the struct in
	fmt.Printf("%+v\n",you)
}