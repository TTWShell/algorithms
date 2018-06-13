package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
