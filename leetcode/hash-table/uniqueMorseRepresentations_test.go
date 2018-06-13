package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
