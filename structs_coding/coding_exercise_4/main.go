/*
1. Create a new struct type called grades with  2 fields: grade and course

2. Add another field of type grades to person struct type (embedded struct).

3. Change the initialization of the struct values called me and you to contain information about the grades.

4. Change the grade and the course of one struct value to "Golang" and 98.

5. Print out the struct values.


*/
package main

import "fmt"

type  person struct {
	name string
	age int
	colors []string	
	gd grades

}

type grades struct {
	grade int
	course string
}
func main(){
	
	me := person{name:"vikas",age:33,colors:[]string{"Red","Green"},gd:grades{grade:9,course:"C++"}}
	you := person{name:"Udemy",age:23,colors:[]string{"Black","White"},gd:grades{grade:8,course:"Python"}}

	fmt.Printf("%v\n",me)
	fmt.Printf("%+v\n",you)
	you.gd.grade = 98
	you.gd.course = "Golang"
	fmt.Printf("%+v\n",you)
}