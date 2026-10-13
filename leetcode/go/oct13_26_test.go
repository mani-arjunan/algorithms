package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxAlternatingSum(t *testing.T) {
	assert.Equal(t, int64(12), maxAlternatingSum([]int{1, 2, 3}))
	assert.Equal(t, int64(16), maxAlternatingSum([]int{1, -1, 2, -2, 3, -3}))
}
