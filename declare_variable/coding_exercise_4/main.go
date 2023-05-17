package main

import "fmt"
// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"
var version = "3.1"
 
func main() {
	name := "Golang"
	fmt.Println(name)
}