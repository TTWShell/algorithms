package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCousins(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 4}, Val: 2}, Val: 1, Right: &TreeNode{Val: 3}}
	assert.False(isCousins(root, 4, 3))

	root = &TreeNode{Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Val: 1, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	assert.True(isCousins(root, 5, 4))

	root = &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 4}, Val: 2}, Val: 1, Right: &TreeNode{Val: 3}}
	assert.False(isCousins(root, 2, 3))
}
