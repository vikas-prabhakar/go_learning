package main

import "fmt"

func main(){
	const secPerDay  = 60 * 60 * 24
	const daysYear = 365
	fmt.Printf("total number of seconds in a year: %d\n",secPerDay * daysYear)

}