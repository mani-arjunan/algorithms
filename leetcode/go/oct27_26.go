package main

import "strings"

func createGrid(m int, n int) []string {
	var result []string

	for m > 0 {
		str := "."
		if m == 1 {
			str += strings.Repeat(".", n-1)
		} else {
			str += strings.Repeat("#", n-1)
		}
		result = append(result, str)
		m--
	}

	return result
}
