package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mostCommonWord(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ball", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
