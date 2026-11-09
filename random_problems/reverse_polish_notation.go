package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []string

	for _, val := range tokens {
		if val == "*" || val == "+" || val == "-" || val == "/" {
			val2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			val1, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			if val == "+" {
				val3 := val1 + val2
				stack = append(stack, strconv.Itoa(val3))
			}
			if val == "*" {
				val3 := val1 * val2
				stack = append(stack, strconv.Itoa(val3))
			}
			if val == "-" {
				val3 := val1 - val2
				stack = append(stack, strconv.Itoa(val3))
			}
			if val == "/" {
				val3 := val1 / val2
				stack = append(stack, strconv.Itoa(val3))
			}
		} else {
			stack = append(stack, val)
		}
	}

	val, _ := strconv.Atoi(stack[0])
	fmt.Println(stack, val)
	return val
}
