package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_toGoatLatin(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Imaa peaksmaaa oatGmaaaa atinLmaaaaa", toGoatLatin("I speak Goat Latin"))
	assert.Equal("heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		toGoatLatin("The quick brown fox jumped over the lazy dog"))
}
