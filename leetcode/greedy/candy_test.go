package lgreedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, candy([]int{1}))
	assert.Equal(5, candy([]int{1, 0, 2}))
	assert.Equal(4, candy([]int{1, 2, 2}))
	assert.Equal(8, candy([]int{1, 0, 2, 0, 1}))
	assert.Equal(13, candy([]int{1, 2, 87, 87, 87, 2, 1}))
	assert.Equal(11, candy([]int{1, 2, 3, 4, 1}))
}
