package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{3, 6}}, largeGroupPositions("abbxxxxzzy"))
	assert.Equal([][]int{}, largeGroupPositions("abc"))
	assert.Equal([][]int{{3, 5}, {6, 9}, {12, 14}}, largeGroupPositions("abcdddeeeeaabbbcd"))
	assert.Equal([][]int{{3, 6}}, largeGroupPositions("abbxxxx"))
}
