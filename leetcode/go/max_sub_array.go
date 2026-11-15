package main

func maxSubArraySum(nums []int) int {
	currentSum := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		temp := currentSum + nums[i]

		if nums[i] > temp {
			currentSum = nums[i]
		} else {
			currentSum = temp
		}

		if currentSum > currentMax {
			currentMax = currentSum
		}
	}

	return currentMax
}
