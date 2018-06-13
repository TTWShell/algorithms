package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestWord(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("world", longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	assert.Equal("apple", longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	assert.Equal("breakfast", longestWord([]string{
		"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast",
		"l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"}))
}
