package main

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		midPoint := low + (high-low)/2
		if nums[midPoint] == target {
			return midPoint
		} else if nums[midPoint] > target {
			high = midPoint - 1
		} else {
			low = midPoint + 1
		}
	}

	return -1
}

func binarySearchRecurse(nums []int, target int) int {
	midPoint := len(nums) / 2

	if nums[midPoint] == target {
		return midPoint
	} else if nums[midPoint] > target {
		return binarySearch(nums[:midPoint], target)
	} else {
		val := binarySearch(nums[midPoint:], target)
		if val != -1 {
			return val + midPoint + 1
		}
	}

	return -1
}

func main() {
	arr := []int{-1, 0, 2, 4, 6, 8}
	binarySearch(arr, 4)
	binarySearch(arr, 4)
}
