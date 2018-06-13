package ldc

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	assert := assert.New(t)

	res := diffWaysToCompute("2-1-1")
	sort.Ints(res)
	assert.Equal([]int{0, 2}, res)

	res = diffWaysToCompute("1+1+1")
	sort.Ints(res)
	assert.Equal([]int{3, 3}, res)

	res = diffWaysToCompute("2*3-4*5")
	sort.Ints(res)
	assert.Equal([]int{-34, -14, -10, -10, 10}, res)
}
