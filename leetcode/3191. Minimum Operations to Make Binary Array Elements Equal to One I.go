/*
You are given a binary array nums.

You can do the following operation on the array any number of times (possibly zero):

Choose any 3 consecutive elements from the array and flip all of them.
Flipping an element means changing its value from 0 to 1, and from 1 to 0.

Return the minimum number of operations required to make all elements in nums equal to 1. If it is impossible, return -1.

 

Example 1:

Input: nums = [0,1,1,1,0,0]

Output: 3

Explanation:
We can do the following operations:

Choose the elements at indices 0, 1 and 2. The resulting array is nums = [1,0,0,1,0,0].
Choose the elements at indices 1, 2 and 3. The resulting array is nums = [1,1,1,0,0,0].
Choose the elements at indices 3, 4 and 5. The resulting array is nums = [1,1,1,1,1,1].
Example 2:

Input: nums = [0,1,1,1]

Output: -1

Explanation:
It is impossible to make all elements equal to 1.

 

Constraints:

3 <= nums.length <= 105
0 <= nums[i] <= 1
*/
func flip(nums []int, i int){
    nums[i] ^= 1
    nums[i + 1] ^= 1
    nums[i + 2] ^= 1
}
func minOperations(nums []int) int {   
  
    // Edge case: If there are fewer than 3 elements, flipping can't be done
    if len(nums) < 3 {
        return -1
    }
  
    count:=0

    // Iterate up to len(nums) - 3 so we don't go out of bounds
    for i := 0; i <= len(nums) - 3; i++{
        
        // If nums[i] is not 1, flip the bits starting at index i
        if nums[i]!=1{
            flip(nums,i)
            count ++
        }
        
    }
    
    // After the loop, check if the last two elements are 1 or not
    if nums[len(nums) -1] == 0 || nums[len(nums) -2] == 0 {
        return -1
    }
    return count
    
}

/* 
TC=O(3*n) as fliping element 3times when current element is not 1 =O(n)
SC=O(1)*/
