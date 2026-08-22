package main

import "fmt"

func main() {
	mainArr := [][]int{
		{5, 2, 1, 6, 4, 3},
		{5, 2, 1, 6, 4, 3},
	}
	i := 1

	for _, arr := range mainArr {
		for i < len(arr) {
			key := arr[i]
			j := i - 1

			for j >= 0 && arr[j] < key {
				arr[j+1] = arr[j]
				arr[j] = key
				j--
			}
			i++
		}

		i = 1
	}

	fmt.Println(mainArr)
}
