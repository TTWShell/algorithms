package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyHashSet(t *testing.T) {
	assert := assert.New(t)

	obj := Constructor()
	obj.Add(1)
	assert.True(obj.Contains(1))
	obj.Add(2)
	assert.True(obj.Contains(2))
	obj.Remove(1)
	assert.False(obj.Contains(1))
	obj.Remove(2)
	assert.False(obj.Contains(2))
}
