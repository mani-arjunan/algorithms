package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack [][]int

	for i, val := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1][1] < val {
			stackIndex := stack[len(stack)-1][]
			stack = stack[:len(stack)-1]
			result[stackIndex] = i - stackIndex
		}

		stack = append(stack, []int{i, val})
	}

	return result
}
