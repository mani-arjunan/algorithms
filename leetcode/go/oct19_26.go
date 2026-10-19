package main

import "sort"

func targetIndices(nums []int, target int) []int {
	result := make([]int, 0)

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	for i, val := range nums {
		if val == target {
			result = append(result, i)
		}
	}

	return result
}
