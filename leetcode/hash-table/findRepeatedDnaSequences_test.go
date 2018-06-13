package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"AAAAACCCCC", "CCCCCAAAAA"}, findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	assert.Equal([]string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAA"))
	assert.Equal([]string{}, findRepeatedDnaSequences("A"))
}
