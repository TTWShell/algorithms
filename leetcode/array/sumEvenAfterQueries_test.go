package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{8, 6, 2, 4}, sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
	assert.Equal([]int{4, 10, 10}, sumEvenAfterQueries([]int{5, 7, 4}, [][]int{{0, 1}, {1, 0}, {4, 1}}))
}
