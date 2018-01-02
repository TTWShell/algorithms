package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	assert := assert.New(t)

	obj := WDConstructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	assert.False(obj.Search(""))
	assert.False(obj.Search("pad"))
	assert.True(obj.Search("bad"))
	assert.True(obj.Search(".ad"))
	assert.False(obj.Search("b."))
	assert.True(obj.Search("b.."))
	assert.True(obj.Search("b.d"))
	assert.False(obj.Search("b.a"))
}
