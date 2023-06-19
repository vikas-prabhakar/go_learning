/*
Rename the file created at Exercise #1 to data.txt

Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
*/

package main
import (
	"os"
	"log"
)
func main() {
	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File doesn't exist")
		}
		
	} 
	err = os.Rename("info.txt","data.txt")
	if err != nil{
		log.Fatal(err)
	} 

}