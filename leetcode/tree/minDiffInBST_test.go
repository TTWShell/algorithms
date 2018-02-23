package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}
	assert.Equal(1, minDiffInBST(root))

	root = &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 100}}}
	assert.Equal(1, minDiffInBST(root))

	root = &TreeNode{Val: 99, Left: &TreeNode{Val: 84, Left: &TreeNode{Val: 27, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 53}}}}
	assert.Equal(15, minDiffInBST(root))
}
