/*
Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at the command line.

The user should give between 2 and 10 numbers.

Notes:

- the program should be run in a terminal (go run main.go) not in Go Playground

- example:

go run main.go 3 2 5

Expected output: Sum: 10, Product: 30

*/
package main
 
import (
    "fmt"
    "os"
    "log"
    "strconv"
)

func main(){
    if len(os.Args) < 2 {
        log.Fatal("Please enter arguments between 2 and 10 numbers")
    }


    sum, product := 0, 1
    if len(os.Args) < 2 &&  len(os.Args) > 10{
        fmt.Printf("Please enter between 2 and 10 numbers")
    } else {
            for _, v := range os.Args[1:]{

                val, _ := strconv.Atoi(v)
                sum += val
                product *= val
    
            }
            fmt.Printf("Sum :%v ,Product :%v\n",sum,product)
        }


}