package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}, shortestToChar("loveleetcode", 'e'))
	assert.Equal([]int{0, 1, 2, 3}, shortestToChar("baaa", 'b'))
	assert.Equal([]int{0, 1, 1, 0}, shortestToChar("baab", 'b'))
	assert.Equal([]int{3, 2, 1, 0}, shortestToChar("aaab", 'b'))
}
