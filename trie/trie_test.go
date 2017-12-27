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

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	T := Constructor()
	T.Insert("a")
	T.Insert("b")
	T.Insert("an")
	T.Insert("android")

	assert.False(T.Search(""))
	assert.False(T.Search("z"))
	assert.True(T.Search("a"))
	assert.True(T.Search("b"))
	assert.True(T.Search("an"))
	assert.False(T.Search("and"))
	assert.True(T.Search("android"))
	assert.False(T.Search("androidfuck"))
}

func TestStartsWith(t *testing.T) {
	assert := assert.New(t)

	T := Constructor()
	T.Insert("a")
	T.Insert("b")
	T.Insert("an")
	T.Insert("android")

	assert.False(T.StartsWith(""))
	assert.False(T.StartsWith("z"))
	assert.True(T.StartsWith("a"))
	assert.True(T.StartsWith("an"))
	assert.True(T.StartsWith("androi"))
	assert.False(T.StartsWith("androidfuck"))
}
