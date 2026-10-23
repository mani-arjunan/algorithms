package main

func maxArea(heights []int) int {
	var result int
	leftPtr, rightPtr := 0, len(heights)

	for leftPtr < rightPtr {
		leftValue := heights[leftPtr]
		rightValue := heights[rightPtr]
		minValue := min(leftValue, rightValue)

		area := (rightPtr - leftPtr) * minValue

		if area > result {
			result = area
		}

		if leftValue < rightValue {
			leftPtr++
		} else {
			rightPtr--
		}

	}

	return result
}
