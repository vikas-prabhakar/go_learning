/*
Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.

The file exists in the current working directory.
*/

package main
 
import (
	"os"
	"bufio"
	"fmt"
	"log"
)
 
func main() {
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err :=scanner.Err(); err != nil {
		log.Fatal(err)
	}
}