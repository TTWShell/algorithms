package larray

import "testing"
import "github.com/stretchr/testify/assert"

func Test_largestPerimeter(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, largestPerimeter([]int{2, 1, 2}))
	assert.Equal(0, largestPerimeter([]int{1, 2, 1}))
	assert.Equal(10, largestPerimeter([]int{3, 2, 3, 4}))
	assert.Equal(8, largestPerimeter([]int{3, 6, 2, 3}))
}
