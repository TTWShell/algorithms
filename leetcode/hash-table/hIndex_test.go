package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hIndex(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, hIndex([]int{3, 0, 6, 1, 5}))
}
