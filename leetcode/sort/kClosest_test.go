package lsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCousins(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{-2, 2}}, kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	assert.Equal([][]int{{3, 3}, {-2, 4}}, kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
