package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParity(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 4, 1, 3}, sortArrayByParity([]int{3, 1, 2, 4}))
}
