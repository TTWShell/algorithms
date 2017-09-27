package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge2(t *testing.T) {
	assert := assert.New(t)

	input := []Interval{Interval{1, 3}, Interval{2, 6}, Interval{8, 10}, Interval{15, 18}}
	assert.Equal(merge2(input), []Interval{Interval{1, 6}, Interval{8, 10}, Interval{15, 18}})
	assert.Equal(merge2([]Interval{}), []Interval{})
}
