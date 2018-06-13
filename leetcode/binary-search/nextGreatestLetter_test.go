package lbs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(byte('c'), nextGreatestLetter([]byte("cfj"), byte('a')))
	assert.Equal(byte('f'), nextGreatestLetter([]byte("cfj"), byte('c')))
	assert.Equal(byte('f'), nextGreatestLetter([]byte("cfj"), byte('d')))
	assert.Equal(byte('j'), nextGreatestLetter([]byte("cfj"), byte('g')))
	assert.Equal(byte('c'), nextGreatestLetter([]byte("cfj"), byte('j')))
	assert.Equal(byte('c'), nextGreatestLetter([]byte("cfj"), byte('k')))

	assert.Equal(byte('f'), nextGreatestLetter([]byte("cf"), byte('d')))

	assert.Equal(byte('v'), nextGreatestLetter([]byte("eeekqqqvvy"), byte('q')))
}
