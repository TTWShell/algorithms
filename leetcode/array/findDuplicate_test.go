package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, findDuplicate([]int{3, 1, 2, 6, 5, 4, 3}))
}
