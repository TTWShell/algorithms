package lbst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}

	assert.Equal(root, searchBST(root, 4))
	assert.Equal(root.Left, searchBST(root, 2))
	assert.Equal(root.Right, searchBST(root, 7))
	assert.Equal(root.Left.Left, searchBST(root, 1))
	assert.Equal(root.Left.Right, searchBST(root, 3))
}
