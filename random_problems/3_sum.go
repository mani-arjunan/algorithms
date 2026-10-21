package main

import "sort"

func threeSum(nums []int) [][]int {
	hashMap := map[[3]int]bool{}
	var result [][]int
	sort.Slice(nums, func(a, b int) bool {
		return nums[b] > nums[a]
	})

	for i, val := range nums {
		left, right := i+1, len(nums)-1

		for left < len(nums) && right >= 0 {
			sum := val + nums[left] + nums[right]

			if sum < 0 {
				left++
				continue
			}

			if sum > 0 {
				right--
				continue
			}

			hashMap[[3]int{val, nums[left], nums[right]}] = true
			left++
			right--
		}

	}

	for key, _ := range hashMap {
		result = append(result, []int{key[0], key[1], key[2]})
	}

	return result
}
