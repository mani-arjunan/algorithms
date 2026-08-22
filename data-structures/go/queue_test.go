package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestBaseCases(t *testing.T) {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	assert.Equal(t, 3, q.length)
	assert.Equal(t, 1, q.dequeue())
	assert.Equal(t, 2, q.head.value)
	assert.Equal(t, 3, q.tail.value)
	assert.Equal(t, 2, q.length)
}

func TestDequeueSingleNode(t *testing.T) {
	q := queue{}
	q.enqueue(1)

	assert.Equal(t, 1, q.length)
	assert.Equal(t, 1, q.dequeue())
	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)
	assert.Equal(t, 0, q.length)
}

func TestEmptyDequeue(t *testing.T) {
	q := queue{}

	assert.Equal(t, 0, q.length)
	assert.Equal(t, -1, q.dequeue())
	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)
}
