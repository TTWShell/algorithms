package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, sumNumbers(nil))
	assert.Equal(25, sumNumbers(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	assert.Equal(137, sumNumbers(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}))
}
