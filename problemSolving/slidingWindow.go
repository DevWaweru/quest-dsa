package problemSolving

import (
	"fmt"
)

// Problem: Find sum of 3 consecutive numbers in an array
// [1,2,3,4,5,6], 3
func findSumSliding(arr []int, k int) {
	var firstSum = 0
	// first get the sum of the first 3 numbers
	for i := 0; i < k; i++ {
		firstSum += arr[i]
	}
	var results = []int{firstSum}
	// get the sum of the rest
	for j := 1; j <= k; j++ {
		firstSum = firstSum - arr[j-1]
		firstSum = firstSum + arr[j+k-1]
		results = append(results, firstSum)
	}
	fmt.Println(results)
}

// Problem:
// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
// 7 [2,3,1,2,4,3] == 2 (4,3)
// 4 [1,4,4] == 1 (1)
// 11 [1,1,1,1,1,1,1,1] == 0
func minimumSubArr(arr []int, target int) {
	// assign the largest number to a variable
	min_val := len(arr)
	// know where to start, where to end, and summation
	var start = 0
	var end = 0
	var summ = 0
	// loop through the entire array; while the end is less than length of arr
	for end < len(arr) {
		// find the sum of the next value
		summ += arr[end]
		end += 1
		// Start removing previous values, if the sum of sub array is more than or equat to target
		for start < end && summ >= target {
			summ -= arr[start]
			start += 1
			// Get the length of the sub array
			if min_val > end-start+1 {
				min_val = end - start + 1
			}
		}
	}
	if min_val == len(arr) {
		min_val = 0
	}
	fmt.Println(min_val)
}

func SlidingWindows() {
	fmt.Println("---- Sliding ----")
	// This is an example
	findSumSliding([]int{1, 2, 3, 4, 5, 6}, 3)
	minimumSubArr([]int{1, 1, 1, 1, 1, 1, 1, 1}, 11)
	fmt.Println("---- -------- ----")
}
