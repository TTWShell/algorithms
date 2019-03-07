package lht

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonChars(t *testing.T) {
	assert := assert.New(t)

	assert.Subset([]string{"e", "l", "l"}, commonChars([]string{"bella", "label", "roller"}))
	assert.Subset(commonChars([]string{"bella", "label", "roller"}), []string{"e", "l", "l"})
	assert.Subset([]string{"c", "o"}, commonChars([]string{"cool", "lock", "cook"}))
	assert.Subset(commonChars([]string{"cool", "lock", "cook"}), []string{"c", "o"})
}
