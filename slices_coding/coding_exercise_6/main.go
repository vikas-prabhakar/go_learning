/*
Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

*/
package main

import "fmt"

func main(){
    friends := []string{"Marry", "John", "Paul", "Diana"}
    copyFriends := make([]string, len(friends))
    copy(copyFriends,friends[0:])
    copyFriends[0]="Peter Parker"
    fmt.Printf("friends : %v \n",friends)
    fmt.Printf("copy of the slice : %v \n",copyFriends)
}