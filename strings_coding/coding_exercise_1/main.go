/*
1. Using the var keyword declare a string called name and initialize it with your name.

2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.

3. Print the following string on multiple lines like this:

Your name: `here the value of name variable`

Country: `here the value of country variable`

4. Print out the following strings:

a) He says: "Hello"

b) C:\Users\a.txt
*/

package main

import (
	"fmt"

)

func main() {
	var name string = "Vikas"
	country := "India"
	fmt.Printf("Your name: %s\nCountry: %s\n",name,country) 
	fmt.Printf("Your name: %s\nCountry: %s\n",name,country) 
	fmt.Printf("Your name: %s\nCountry: %s\n",name,country) 
	fmt.Printf("Your name: %s\nCountry: %s\n",name,country) 
	fmt.Printf("He says: \"Hello\"\n")
	fmt.Printf("C:\\Users\\a.txt")
}
