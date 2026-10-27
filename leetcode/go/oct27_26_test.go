package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestCreateGrid(t *testing.T) {
	assert.Equal(t, []string{"...."}, createGrid(1, 4))
	assert.Equal(t, []string{".##", ".##", "..."}, createGrid(3, 3))
	assert.Equal(t, []string{".##", "..."}, createGrid(2, 3))
}
