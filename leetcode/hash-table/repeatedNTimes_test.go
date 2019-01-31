package lht

import "testing"
import "github.com/stretchr/testify/assert"

func Test_repeatedNTimes(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, repeatedNTimes([]int{1, 2, 3, 3}))
	assert.Equal(2, repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	assert.Equal(5, repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}
