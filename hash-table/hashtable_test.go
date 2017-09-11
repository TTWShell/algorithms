package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
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
	assert.Equal(ht.count, 1)

	// test update exist key
	ht.Put(1, 235423423)
	assert.Equal(ht.count, 1)
	assert.True(ht.items[ht.hash(1)].value == 235423423)

	// test conflict management
	assert.Equal(ht.hash(1), ht.hash(14))
	ht.Put(14, 12233)
	assert.Equal(ht.count, 2)
	assert.True(ht.items[ht.hash(1)].key == 1)
	assert.True(ht.items[ht.hash(1)+1].key == 14, "Because use linearProbing")

	defer func() {
		assert.Equal(ht.count, MinSize)
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
	assert.Equal(ht.count, 1)
	value, err := ht.Pop(1)
	assert.True(value == 2)
	assert.Equal(ht.count, 0)

	value, err = ht.Pop(1)
	assert.Equal(ht.count, 0)
	assert.True(err.Error() == "key no exists")
}

func TestMix(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)
	ht.Put(1, 22313231)
	ht.Put("a", []int{1, 2, 3})

	value, err := ht.Get(1)
	assert.True(value == 22313231)

	value, err = ht.Pop("a")
	assert.True(reflect.DeepEqual(value, []int{1, 2, 3}))

	value, err = ht.Pop("1")
	assert.True(err.Error() == "key no exists")

	defer func() {
		assert.Equal(recover(), "Only support int、string type key")
	}()
	ht.Put('a', []int{1, 2, 3})
}

func Test_rehash(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize)

	ht.Put(1, 22313231)
	ht.Put("abc", []int{1, 2, 3})

	ht.rehash()
	assert.Equal(ht.size, MinSize*2)
	assert.Equal(len(ht.items), MinSize*2)

	value, _ := ht.Get(1)
	assert.True(value == 22313231)

	value, _ = ht.Get("abc")
	assert.True(reflect.DeepEqual(value, []int{1, 2, 3}))
}
