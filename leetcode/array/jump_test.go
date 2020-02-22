package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, jump([]int{2, 3, 1, 1, 4}))
	// 按题意不会出现，greed方法下面返回5，实际不可达
	// assert.Equal(0, jump([]int{2, 3, 1, 1, 0, 4}))
	assert.Equal(1, jump([]int{2, 1}))

	assert.Equal(0, jump([]int{1}))
}
