/*
Remove the file created at exercise #1. Take care that the file is now called data.txt (it has been renamed at exercise #2).

Perform error handling.
*/

package main

import (
	"os"
	"log"
)

func main() {
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
