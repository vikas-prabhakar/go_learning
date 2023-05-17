package main
 
import "fmt" 
func main() {
	const (
		Jun = iota + 6
		Jul 
		Aug 

	)

    fmt.Println("Jun :",Jun)
	fmt.Println("Jul :",Jul)
	fmt.Println("Aug :",Aug)
}