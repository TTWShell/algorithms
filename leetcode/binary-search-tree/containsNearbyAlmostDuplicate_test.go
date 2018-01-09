package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	assert := assert.New(t)

	assert.True(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, 0))
	assert.False(containsNearbyAlmostDuplicate([]int{}, 0, 0))
}
