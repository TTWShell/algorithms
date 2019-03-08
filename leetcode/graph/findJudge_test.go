package lgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findJudge(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, findJudge(1, [][]int{}))
	assert.Equal(2, findJudge(2, [][]int{{1, 2}}))
	assert.Equal(3, findJudge(3, [][]int{{1, 3}, {2, 3}}))
	assert.Equal(-1, findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	assert.Equal(-1, findJudge(3, [][]int{{1, 2}, {2, 3}}))
	assert.Equal(3, findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}))
}
