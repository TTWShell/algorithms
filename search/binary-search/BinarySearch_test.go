package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	assert := assert.New(t)
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
	for index, target := range sorted {
		assert.Equal(BinarySearch(sorted, target), index)
	}

	assert.Equal(BinarySearch(sorted, 10), -1)

	assert.Equal(BinarySearch(sorted, 16), -1)
	assert.Equal(BinarySearch(sorted, 0), -1)
}
