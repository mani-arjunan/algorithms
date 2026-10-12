package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestRemoveOuterParenthesis(t *testing.T) {
	str := map[string]string{
		"(()())(())":         "()()()",
		"(()())(())(()(()))": "()()()()(())",
		"()()":               "",
	}

	for key, val := range str {
		result := removeOuterParanthesis(key)

		assert.Equal(t, val, result)
	}
}
