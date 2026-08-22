package main

import "fmt"

func merge(left, right []int) []int {
	var result []int
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i <= len(left) {
		result = append(result, left[i:]...)
	}
	if j <= len(right) {
		result = append(result, right[j:]...)
	}
	fmt.Println(result, i, j, left, right)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	midPoint := len(arr) / 2

	left := mergeSort(arr[0:midPoint])
	right := mergeSort(arr[midPoint:])

	finalArr := merge(left, right)
	return finalArr
}

func main() {
	arr := []int{2, 4, 5, 7, 1, 3, 6}
	finalArr := mergeSort(arr)

	fmt.Println(finalArr)
}
