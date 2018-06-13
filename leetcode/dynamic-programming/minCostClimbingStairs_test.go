package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(15, minCostClimbingStairs([]int{10, 15, 20}))
	assert.Equal(6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	assert.Equal(0, minCostClimbingStairs([]int{0, 0, 0, 1}))
	assert.Equal(1, minCostClimbingStairs([]int{0, 1, 1, 0}))
}
