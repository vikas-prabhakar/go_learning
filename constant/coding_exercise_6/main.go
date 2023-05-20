/*

Using Iota declare the following months of the year: Jun, Jul and Aug

Jun, Jul and Aug are constant and their value is 6, 7 and 8.
*/

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