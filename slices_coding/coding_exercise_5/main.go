/*
Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.

Print those elements and their sum.

*/
package main

import "fmt"

func main(){
    nums := []int{5, -1, 9, 10, 1100, 6, -1, 6} 
    sum :=0 
    for _, v := range nums[1: len(nums)-2]{
        sum += v
    }
    fmt.Printf("Elements : %v and their Sum: %d\n",nums[1: len(nums)-2], sum)
}