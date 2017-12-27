package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	T := Constructor()

	T.Insert("a")
	assert.Len(T.data, 1)
	assert.True(T.data['a'].IsEnd)
	assert.Equal(T.data['a'].Val, 'a')
	assert.Empty(T.data['a'].Next)

	T.Insert("b")
	assert.Len(T.data, 2)
	assert.True(T.data['b'].IsEnd)
	assert.Equal(T.data['b'].Val, 'b')
	assert.Empty(T.data['b'].Next)

	T.Insert("an")
	assert.Len(T.data, 2)
	assert.True(T.data['a'].IsEnd)
	assert.Equal(T.data['a'].Val, 'a')
	assert.Len(T.data['a'].Next, 1)
	assert.True(T.data['a'].Next['n'].IsEnd)
	assert.Equal(T.data['a'].Next['n'].Val, 'n')
	assert.Empty(T.data['a'].Next['n'].Next)

	defer func() {
		if r := recover(); r != "Length of word must > 0." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	// test panic
	T.Insert("")
}
