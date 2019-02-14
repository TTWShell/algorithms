package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 0}, Val: 2, Right: &TreeNode{Val: 4}}, Val: 6, Right: &TreeNode{Val: 8}}
	assert.Equal(6, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 8}).Val)
	assert.Equal(2, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 4}).Val)
}
