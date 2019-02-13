package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyHashMap(t *testing.T) {
	assert := assert.New(t)

	hashMap := HashMapConstructor()

	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	assert.Equal(1, hashMap.Get(1))
	assert.Equal(-1, hashMap.Get(3))
	hashMap.Put(2, 1)
	assert.Equal(1, hashMap.Get(2))
	hashMap.Remove(2)
	assert.Equal(-1, hashMap.Get(2))
}
