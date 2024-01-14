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
	arrMap := make(map[int]int)
	fmt.Println(arrMap)
	for i, num := range arr {
		diff := target - num

		if index, ok := arrMap[diff]; ok {
			fmt.Println([]int{index, i})
			return []int{index, i}
		}
		arrMap[num] = i
	}
	return []int{}
}

func TwoPointers() {
	twoSum([]int{3, 2, 4}, 6)
}
