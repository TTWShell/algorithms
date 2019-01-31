package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDistToClosest(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	assert.Equal(3, maxDistToClosest([]int{1, 0, 0, 0}))
	assert.Equal(3, maxDistToClosest([]int{0, 0, 0, 1}))
}
