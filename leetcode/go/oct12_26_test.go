package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSumOfSquares(t *testing.T) {
	assert.Equal(t, 21, sumOfSquares([]int{1, 2, 3, 4}))
	assert.Equal(t, 63, sumOfSquares([]int{2, 7, 1, 19, 18, 3}))
}
