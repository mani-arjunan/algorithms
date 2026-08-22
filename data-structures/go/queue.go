package main

import "fmt"

type node struct {
	value int
	next  *node
}

type queue struct {
	length int
	head   *node
	tail   *node
}

func (q *queue) enqueue(value int) {
	n := node{
		value: value,
	}

	q.length += 1
	if q.head == nil {
		q.head = &n
		q.tail = &n
		return
	}

	q.tail.next = &n
	q.tail = &n
}

func (q *queue) dequeue() int {
	if q.head == nil {
		return -1
	}

	q.length -= 1
	val := q.head.value

	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return val
	}
	q.head = q.head.next
	return val
}

func testQueue() {
	q := queue{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.dequeue()
	q.dequeue()
	q.dequeue()

	fmt.Println(q.head)
	fmt.Println(q.tail)
}
