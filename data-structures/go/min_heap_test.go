package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestBaseCasesForMinHeap(t *testing.T) {
	minHeap := Heap{}

	minHeap.Insert(5)
	minHeap.Insert(6)
	minHeap.Insert(1)

	assert.Equal(t, 3, minHeap.length)
	assert.Equal(t, 1, minHeap.Delete())
	assert.Equal(t, 2, minHeap.length)
}

func TestGettingEmptyMinHeap(t *testing.T) {
	minHeap := Heap{}

	assert.Equal(t, -1, minHeap.Delete())
	assert.Equal(t, 0, minHeap.length)
}
