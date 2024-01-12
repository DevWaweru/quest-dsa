package problemSolving

import "fmt"

// Problem:
// Given an array of integers nums and an integer target, return
// indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// [2,7,11,15], 9 => [0,1]
// [3,2,4], 6 => [1,2]
// [3,3], 6 => [0,1]
func twoSum(arr []int, target int) []int {
	var start = 0
	var end = len(arr) - 1
	for start < end {
		summ := arr[start] + arr[end]
		if summ == target {
			fmt.Println([]int{start, end})
			return []int{start, end}
		} else if summ <= target {
			start += 1
		} else {
			end -= 1
		}
	}
	fmt.Println(arr)
	return arr
}

func TwoPointers() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
