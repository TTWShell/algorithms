package lht

import "testing"
import "github.com/stretchr/testify/assert"

func Test_uncommonFromSentences(t *testing.T) {
	assert := assert.New(t)

	assert.Subset([]string{"sweet", "sour"}, uncommonFromSentences("this apple is sweet", "this apple is sour"))
	assert.Equal([]string{"banana"}, uncommonFromSentences("apple apple", "banana"))
}
