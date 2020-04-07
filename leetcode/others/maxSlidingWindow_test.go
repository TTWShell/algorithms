package lothers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSlidingWindow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal([]int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
}
