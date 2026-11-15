package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxSubArraySum(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	input2 := []int{1}
	input3 := []int{5, 4, -1, 7, 8}

	assert.Equal(t, 6, maxSubArraySum(input))
	assert.Equal(t, 1, maxSubArraySum(input2))
	assert.Equal(t, 23, maxSubArraySum(input3))
}
