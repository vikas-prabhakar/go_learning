/*
Create a Go Program that:

1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.

2. Write the string "The Go gopher is an iconic mascot!" to the file.
*/

package main
 
import (
	"os"
	"log"

)
 
func main() {
	file, err := os.OpenFile("info.txt",os.O_WRONLY,0644)
	if err != nil {
		log.Fatal(err)
	}
	data := []byte("The Go gopher is an iconic mascot!")
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}