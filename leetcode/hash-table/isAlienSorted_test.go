package lht

import "testing"
import "github.com/stretchr/testify/assert"

func Test_isAlienSorted(t *testing.T) {
	assert := assert.New(t)

	assert.True(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	assert.False(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	assert.False(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
	assert.True(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))
}
