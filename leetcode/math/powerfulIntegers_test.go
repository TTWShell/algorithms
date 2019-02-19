package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_powerfulIntegers(t *testing.T) {
	assert := assert.New(t)

	assert.Subset([]int{9, 17, 33, 65, 2, 3, 5}, powerfulIntegers(1, 2, 100))
	assert.Subset(powerfulIntegers(1, 2, 100), []int{9, 17, 33, 65, 2, 3, 5})

	assert.Subset([]int{2, 3, 4, 5, 7, 9, 10}, powerfulIntegers(2, 3, 10))
	assert.Subset(powerfulIntegers(2, 3, 10), []int{2, 3, 4, 5, 7, 9, 10})

	assert.Subset([]int{2, 4, 6, 8, 10, 14}, powerfulIntegers(3, 5, 15))
	assert.Subset(powerfulIntegers(3, 5, 15), []int{2, 4, 6, 8, 10, 14})
}
