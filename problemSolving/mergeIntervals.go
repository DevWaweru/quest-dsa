package problemSolving

import (
	"fmt"
	"sort"
)

// Problem:
// Given an array of intervals where intervals[i] = [start, end], merge all overlapping intervals
// and return an array of the non overlapping intervals that cover all the intervals in the input
// intervals = [[1,3],[2,6],[8,10],[15,18]], ==> [[1,6],[8,10],[15,18]]
// intervals = [[1,4],[4,5]], ==> [[1,5]]

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func nonOverlapping(intervals [][]int) [][]int {
	const START, END = 0, 1
	result := make([][]int, 0)
	sort.Slice(intervals, func(a, b int) bool {
		fmt.Println(intervals[a][0], intervals[b][0])
		fmt.Println(intervals[a][1], intervals[b][1])
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})

	for _, curInterval := range intervals {
		if (len(result) == 0) || result[len(result)-1][END] < curInterval[START] {
			// no overlapping
			result = append(result, curInterval)
		} else {
			// has overlapping
			// merge with previous interval
			result[len(result)-1][END] = max(result[len(result)-1][END], curInterval[END])
		}
	}
	fmt.Println(result)
	return result
}

func MergeIntervals() {
	inter1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	inter2 := [][]int{{1, 4}, {4, 5}}
	nonOverlapping(inter1)
	nonOverlapping(inter2)
}
