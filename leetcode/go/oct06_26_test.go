package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMinAdditionToMakeValidString(t *testing.T) {
	str := map[string]int{
		"())":      1,
		"(((":      3,
		")))":      3,
		")())(())": 2,
		"(())((":   2,
	}

	for key, val := range str {
		result := minAdditionToMakeValidString(key)

		assert.Equal(t, val, result)
	}
}
