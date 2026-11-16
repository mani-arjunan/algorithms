package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestDigitSum(t *testing.T) {
	assert.Equal(t, "000", digitSum("00000000", 3))
	assert.Equal(t, "135", digitSum("11111222223", 3))
}
