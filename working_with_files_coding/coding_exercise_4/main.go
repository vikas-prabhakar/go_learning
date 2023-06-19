/*
Create a Go Program that reads the entire contents of a file called info.txt into a string. You can use ioutil.ReadAll() or any other function you want.

The file exists in the current working directory.
*/

package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

)
 
func main() {
	file,err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content,err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contents of a file: %s\n",content)

}