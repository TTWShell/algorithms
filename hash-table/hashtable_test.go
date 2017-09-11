package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHT(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)
	assert.Equal(ht.size, MinSize)
	assert.Equal(len(ht.items), MinSize)

	defer func() {
		assert.Equal(recover(), fmt.Sprintf("size must > %d", MinSize))
	}()
	NewHT(MinSize - 1)
}

func Test_hash(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(20)
	for i := 0; i < ht.size; i++ {
		assert.True(ht.hash(i) < ht.size)
	}

	for _, s := range []string{"01234", "2134", "21349", "dawdw", "2123ad", "ada", "212"} {
		assert.True(ht.hash(s) < ht.size)
	}
}

func TestPut(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)
	ht.Put(1, 2)
	assert.True(ht.items[ht.hash(1)] != nil)
	assert.True(ht.items[ht.hash(1)].key == 1)
	assert.True(ht.items[ht.hash(1)].value == 2)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)

	ht.Put(1, 2)
	value, err := ht.Get(1)
	assert.True(value == 2)
	assert.True(err == nil)

	assert.Equal(ht.hash(1), ht.hash(14))
	value, err = ht.Get(14)
	assert.True(err.Error() == "key no exists")
}
