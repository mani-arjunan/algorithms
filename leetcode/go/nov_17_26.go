package main

import (
	"strconv"
)

func digitSum(s string, k int) string {
	copyS := s

	for len(copyS) > k {
		times := (len(copyS) + k - 1) / k
		var temp string

		for i := 0; i < times; i++ {
			start := i * k
			var sub string
			if times-1 == i {
				sub = copyS[start:len(copyS)]
			} else {
				sub = copyS[start:(start + k)]
			}
			var addResult int
			for i := 0; i < len(sub); i++ {
				val, _ := strconv.Atoi(string(sub[i]))
				addResult += val
			}
			temp += strconv.Itoa(addResult)
		}
		copyS = temp
	}

	return copyS
}
