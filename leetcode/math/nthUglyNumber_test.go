package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nthUglyNumber(t *testing.T) {
	assert := assert.New(t)

	for i, num := range []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12} {
		assert.Equal(num, nthUglyNumber(i+1), fmt.Sprintf("%d:%d", i, num))
	}
}

func Benchmark_nthUglyNumber(t *testing.B) {
	assert := assert.New(t)

	assert.Equal(2123366400, nthUglyNumber(1690))
}
