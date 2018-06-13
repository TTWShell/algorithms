package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{}, summaryRanges([]int{}))
	assert.Equal([]string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	assert.Equal([]string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
