package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestTargetIndices(t *testing.T) {
	assert.Equal(t, []int{1, 2}, targetIndices([]int{1, 2, 5, 2, 3}, 2))
	assert.Equal(t, []int{3}, targetIndices([]int{1, 2, 5, 2, 3}, 3))
	assert.Equal(t, []int{4}, targetIndices([]int{1, 2, 5, 2, 3}, 5))
}
