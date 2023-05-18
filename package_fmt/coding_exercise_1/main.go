/*
Consider the following variable declarations: 

x, y, z := 10, 15.5, "Gophers"
score := []int{10, 20, 30}
Using fmt.Printf():

1. Print each variable using a specific verb for its type

2. Print the string value enclosed by double quotes ("Gophers")

3. Print each variable using the same verb

4. Print the type of y and score
*/
package main

import "fmt"

func main(){
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("x using specific verb for its type is: %d\ny using specific verb for its type is:%f\nz using specific verb for its type is: %s\nscore using specific verb for its type is: %#v\n",x,y,z,score)
	fmt.Printf("Print the string value enclosed by double quotes (\"Gophers\")")
	fmt.Printf("z enclosed by double quotes: %q\n",z)
	fmt.Printf("x using the same verb %v\ny using the same verb %v\nz using the same verb %v\nscore using the same verb %v\n",x,y,z,score)
	fmt.Printf("type of y %T and type of score %T\n",y,score)



}