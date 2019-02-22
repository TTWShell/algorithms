package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseOnlyLetters(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("dc-ba", reverseOnlyLetters("ab-cd"))
	assert.Equal("j-Ih-gfE-dCba", reverseOnlyLetters("a-bC-dEf-ghIj"))
	assert.Equal("Qedo1ct-eeLg=ntse-T!", reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
