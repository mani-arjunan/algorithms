package main

import "sort"

func carFleet(target int, position, speed []int) int {
	posSpeed := make([][2]int, len(position))

	for i, val := range position {
		posSpeed[i] = [2]int{val, speed[i]}
	}

	sort.Slice(posSpeed, func(i, j int) bool {
		return posSpeed[i][0] > posSpeed[j][0]
	})

	fleets := 1
	prevTime := float64(target-posSpeed[0][0]) / float64(posSpeed[0][1])

	for i := 1; i < len(posSpeed); i++ {
		curr := float64(target-posSpeed[i][0]) / float64(posSpeed[i][1])

		if prevTime < curr {
			fleets++
			prevTime = curr
		}
	}

	return fleets
}
