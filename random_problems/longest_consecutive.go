package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	var result int
	tempCount := 1
	startValue := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == startValue {
			continue
		}
		if nums[i]-startValue == 1 {
			tempCount++
		} else {
			if result < tempCount {
				result = tempCount
			}
			tempCount = 1
		}
		startValue = nums[i]

	}

	return result
}
