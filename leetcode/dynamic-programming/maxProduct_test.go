package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, maxProduct([]int{}))
	assert.Equal(6, maxProduct([]int{2, 3, -2, 4}))
}
