package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHT(t *testing.T) {
	assert := assert.New(t)
	ht := NewHT(MinSize + 1)
	assert.Equal(ht.size, MinSize+1)

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
