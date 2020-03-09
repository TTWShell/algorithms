package lsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumGap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, maximumGap([]int{3, 6, 9, 1}))
	assert.Equal(0, maximumGap([]int{10}))
}
