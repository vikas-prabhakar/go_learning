/*
Create a new file in the current working directory called info.txt and then close the file. If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
*/

package main

import (
//	"fmt"
	"os"
	"log"
)

func main(){
	 newFile, err := os.Create("info.txt")
	 if err != nil{
		 log.Fatal(err)
	 }
	 newFile.Close()
}
