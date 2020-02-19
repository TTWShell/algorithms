package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, coinChange([]int{1, 2, 5}, 11))
	assert.Equal(-1, coinChange([]int{2}, 3))
	assert.Equal(0, coinChange([]int{1}, 0))
	assert.Equal(2, coinChange([]int{1}, 2))
	assert.Equal(20, coinChange([]int{186, 419, 83, 408}, 6249))
}
