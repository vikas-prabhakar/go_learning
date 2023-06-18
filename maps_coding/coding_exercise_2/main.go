/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main
 
import "fmt"
 
func main() {
	var m1 map[int]bool
	m1[5] = true
 
	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
 
	fmt.Println(m2 == m3)
}
*/

package main
import "fmt"
func main() {
	var m1 map[int]bool
	//m1[5] = true
 
	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
	s1 := fmt.Sprintf("%d", m2)
	s2 := fmt.Sprintf("%d", m3)
	fmt.Println(s1 == s2)
	_ = m1
}