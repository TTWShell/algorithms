package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderLogFiles(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}, reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}
