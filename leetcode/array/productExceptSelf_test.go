package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))
}
