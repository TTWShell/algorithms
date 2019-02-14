package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor2(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 100}, Val: 2, Right: &TreeNode{Val: 4}}, Val: 1, Right: &TreeNode{Val: 8}}
	assert.Equal(1, lowestCommonAncestor2(root, &TreeNode{Val: 2}, &TreeNode{Val: 8}).Val)
	assert.Equal(2, lowestCommonAncestor2(root, &TreeNode{Val: 2}, &TreeNode{Val: 4}).Val)
	assert.Equal(2, lowestCommonAncestor2(root, &TreeNode{Val: 4}, &TreeNode{Val: 2}).Val)
}
