package lht

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uncommonFromSentences(t *testing.T) {
	assert := assert.New(t)

	excepted := []string{"sour", "sweet"}
	result := uncommonFromSentences("this apple is sweet", "this apple is sour")
	sort.Strings(result)
	assert.Equal(excepted, result)
	assert.Equal([]string{"banana"}, uncommonFromSentences("apple apple", "banana"))
}
