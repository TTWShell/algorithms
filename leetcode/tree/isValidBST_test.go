package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	assert := assert.New(t)

	assert.True(isValidBST(&TreeNode{Left: &TreeNode{Val: 1}, Val: 2, Right: &TreeNode{Val: 3}}))
	assert.False(isValidBST(&TreeNode{Left: &TreeNode{Val: 2}, Val: 1, Right: &TreeNode{Val: 3}}))

	assert.False(isValidBST(&TreeNode{Left: &TreeNode{Val: 1}, Val: 2, Right: &TreeNode{Left: &TreeNode{Val: -1}, Val: 3}}))
}
