package main

import "fmt"

type Heap struct {
	data   []int
	length int
}

func (h *Heap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *Heap) Delete() int {
	if h.length == 0 {
		return -1
	}

	headElement := h.data[0]

	if h.length == 1 {
		h.data = []int{}
		h.length = 0
		return headElement
	}

	h.data[0] = h.data[len(h.data)-1]
	h.length--
	h.heapifyDown(0)
	h.data = h.data[:len(h.data)-1]
	return headElement
}

func (h *Heap) getParent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := h.getParent(index)
	parentValue := h.data[parentIndex]
	currentValue := h.data[index]

	if currentValue < parentValue {
		h.data[parentIndex] = currentValue
		h.data[index] = parentValue
		h.heapifyUp(parentIndex)
	}
}

func (h *Heap) getLeftChildIndex(index int) int {
	return (2 * index) + 1
}

func (h *Heap) getRightChildIndex(index int) int {
	return (2 * index) + 2
}

func (h *Heap) heapifyDown(index int) {
	leftChildIndex := h.getLeftChildIndex(index)
	rightChildIndex := h.getRightChildIndex(index)

	if index >= h.length || leftChildIndex >= h.length {
		return
	}

	leftValue := h.data[leftChildIndex]
	currentValue := h.data[index]

	if rightChildIndex >= h.length {
		if currentValue > leftValue {
			h.data[index] = leftValue
			h.data[leftChildIndex] = currentValue
		}
		return
	}

	rightValue := h.data[rightChildIndex]

	if rightValue <= leftValue && currentValue > rightValue {
		h.data[index] = rightValue
		h.data[rightChildIndex] = currentValue
		h.heapifyDown(rightChildIndex)
	} else if rightValue > leftValue && currentValue > leftValue {
		h.data[index] = leftValue
		h.data[leftChildIndex] = currentValue
		h.heapifyDown(leftChildIndex)
	}
}

func testCostMinimization() {
	// arr := []int{9, 8, 2}
	arr := []int{5, 3, 5, 2}
	heap := Heap{}
	total := 0

	for _, val := range arr {
		heap.Insert(val)
	}

	for heap.length > 1 {
		elem1 := heap.Delete()
		elem2 := heap.Delete()
		total += elem1 + elem2
		heap.Insert(elem1 + elem2)
	}

	fmt.Println(total)
}
