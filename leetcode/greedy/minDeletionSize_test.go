package lgreedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDeletionSize(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, minDeletionSize([]string{"cba", "daf", "ghi"}))
	assert.Equal(0, minDeletionSize([]string{"a", "b"}))
	assert.Equal(3, minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
