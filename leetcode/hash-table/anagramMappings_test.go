package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_anagramMappings(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 4, 3, 2, 0}, anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}))
}
