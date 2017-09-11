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

	// test update exist key
	ht.Put(1, 235423423)
	assert.True(ht.items[ht.hash(1)].value == 235423423)

	// test conflict management
	assert.Equal(ht.hash(1), ht.hash(14))
	ht.Put(14, 12233)
	assert.True(ht.items[ht.hash(1)].key == 1)
	assert.True(ht.items[ht.hash(1)+1].key == 14, "Because use linearProbing")

	defer func() {
		assert.Equal(recover(), "hash table overflow")
	}()
	for i := 0; i < 11; i++ {
		ht.Put(i, i)
	}
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

func TestPop(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)

	ht.Put(1, 2)
	value, err := ht.Pop(1)
	assert.True(value == 2)

	value, err = ht.Pop(1)
	assert.True(err.Error() == "key no exists")
}
