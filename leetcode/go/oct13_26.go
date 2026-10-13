package main

import "sort"

func maxAlternatingSum(nums []int) int64 {
	var result int64

	for i, val := range nums {
		nums[i] = val * val
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	midPoint := len(nums) / 2

	if len(nums)%2 == 0 {
		midPoint = midPoint - 1
	}

	for i := 0; i <= midPoint; i++ {
		result += int64(nums[i])
	}

	for i := midPoint + 1; i < len(nums); i++ {
		result -= int64(nums[i])
	}

	return result
}
