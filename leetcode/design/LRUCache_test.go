package ldesign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRUCache(t *testing.T) {
	assert := assert.New(t)

	cache := LRUCacheConstructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(1, cache.Get(1))
	cache.Put(3, 3)                // evicts key 2
	assert.Equal(-1, cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)                // evicts key 1
	assert.Equal(-1, cache.Get(1)) // returns -1 (not found)
	assert.Equal(3, cache.Get(3))  // returns 3
	assert.Equal(4, cache.Get(4))  // returns 4

	cache = LRUCacheConstructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	assert.Equal(-1, cache.Get(1))
	assert.Equal(3, cache.Get(2))
}
