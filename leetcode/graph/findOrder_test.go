package leetcode

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 1}, findOrder(2, [][]int{{1, 0}}))
	res := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	assert.True(reflect.DeepEqual([]int{0, 1, 2, 3}, res) || reflect.DeepEqual([]int{0, 2, 1, 3}, res))
	assert.Equal([]int{2, 1, 0}, findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))

	assert.Equal([]int{}, findOrder(3, [][]int{{0, 1}, {1, 0}}))
}
