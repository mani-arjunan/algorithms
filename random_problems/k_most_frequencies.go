package main

import "fmt"

type Heap struct {
	data   [][]int
	length int
}

func (h *Heap) Insert(value []int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *Heap) Delete() []int {
	if h.length == 0 {
		return []int{}
	}

	headElement := h.data[0]

	if h.length == 1 {
		h.data = [][]int{}
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

	if currentValue[0] > parentValue[0] {
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
		if currentValue[0] < leftValue[0] {
			h.data[index] = leftValue
			h.data[leftChildIndex] = currentValue
		}
		return
	}

	rightValue := h.data[rightChildIndex]

	if rightValue[0] >= leftValue[0] && currentValue[0] < rightValue[0] {
		h.data[index] = rightValue
		h.data[rightChildIndex] = currentValue
		h.heapifyDown(rightChildIndex)
	} else if rightValue[0] < leftValue[0] && currentValue[0] < leftValue[0] {
		h.data[index] = leftValue
		h.data[leftChildIndex] = currentValue
		h.heapifyDown(leftChildIndex)
	}
}


// one of the neetcode 150 problem, i have solved with sorting with hashtable, and it got accepted, and then looked the solution there are ways to solve
// this using heap, so i tried with both minHeap and maxHeap in here.
func main() {
	arr := []int{1, 2, 2, 3, 3, 3}
	// arr := []int{7, 7}
	heap := Heap{}
	k := 2
	count := 0
	var result []int
	hashMap := make(map[int]int)

	for _, val := range arr {
		hashMap[val]++
	}

	for key, val := range hashMap {
		heap.Insert([]int{val, key})
	}

	fmt.Println(heap.data)

	for count < k {
		elem := heap.Delete()
		result = append(result, elem[1])
		count++
	}
	fmt.Println(result)

}
