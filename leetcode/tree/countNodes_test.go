package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countNodes(t *testing.T) {
	assert := assert.New(t)

	var root *TreeNode
	assert.Equal(0, countNodes(root))
	assert.Equal(1, countNodes(&TreeNode{Val: 1}))

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	assert.Equal(3, countNodes(root))
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	assert.Equal(4, countNodes(root))
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	assert.Equal(5, countNodes(root))
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	assert.Equal(6, countNodes(root))
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	assert.Equal(7, countNodes(root))
}
