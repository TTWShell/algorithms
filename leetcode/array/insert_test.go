package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{1, 5}, {6, 9}}, insert([][]int{
		{1, 3}, {6, 9}}, []int{2, 5}))
	assert.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))

	assert.Equal([][]int{{1, 5}}, insert([][]int{}, []int{1, 5}))
	assert.Equal([][]int{{1, 6}}, insert([][]int{{1, 5}}, []int{2, 6}))
	assert.Equal([][]int{{1, 5}, {6, 8}}, insert([][]int{{1, 5}}, []int{6, 8}))
	assert.Equal([][]int{{1, 8}}, insert([][]int{{1, 5}, {6, 8}}, []int{5, 6}))
	assert.Equal([][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	assert.Equal([][]int{{1, 8}}, insert([][]int{{1, 5}, {6, 8}}, []int{3, 7}))

	assert.Equal([][]int{{0, 10}}, insert([][]int{{2, 5}, {6, 7}, {8, 9}}, []int{0, 10}))

	assert.Equal([][]int{{3, 5}, {6, 6}, {12, 15}}, insert([][]int{{3, 5}, {12, 15}}, []int{6, 6}))
}
