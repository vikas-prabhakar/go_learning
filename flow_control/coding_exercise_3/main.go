/*
Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
*/
package main

import 	"fmt"


func main(){
	var count = 0
	for i := 1; i <= 50; i++ {
		if i%7 == 0{
			continue
		}
		fmt.Printf("%d is divisible by 7 \n",i)
		count +=1	
		if count == 3{
			break
		}
	}
}