package main

func sumOfSquares(nums []int) int {
	var total int
	n := len(nums)

	for index, val := range nums {
		if n%(index+1) == 0 {
			total += (val * val)

		}
	}

	return total
}
