package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParityII(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{4, 5, 2, 7}, sortArrayByParityII([]int{4, 2, 5, 7}))
}
