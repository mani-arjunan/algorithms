package main

// Solution1
func twoSum1(nums []int, target int) []int {
	var result []int
	hashMap := make(map[int]int)

	for i, val := range nums {
		diff := target - val
		v, ok := hashMap[diff]

		if ok {
			if v < i {
				result = append(result, v+1, i+1)
			}
			break
		}
		hashMap[val] = i
	}

	return result
}

// Solution2
func twoSum2(nums []int, target int) []int {
	var result []int
	i, j := 0, len(nums)-1

	for i < len(nums) && j >= 0 {
		sum := nums[i] + nums[j]

		if sum == target {
			result = append(result, i+1, j+1)
			break
		}
		if sum >= target {
			j--
		} else {
			i++
		}

	}

	return result
}

func twoSum(nums []int, target int) {
	twoSum1(nums, target)
	twoSum2(nums, target)
}
