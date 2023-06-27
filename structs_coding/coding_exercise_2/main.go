/*
1. Change the name or the struct value called me to "Andrei"

2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.

3. Add a new favorite color to the second struct value called you.

4. Print out the struct values.	
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
	me.name = "Andrei"
	fmt.Printf("%v\n",me)

	colors := you.colors
	fmt.Printf("Type:%T and %v\n",colors,colors)

	colors = append(colors,"Blue")
	you.colors = colors
	
	fmt.Printf("%v\n",me)
	fmt.Printf("%v\n",you)


}