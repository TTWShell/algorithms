package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal([][]int{{3}, {20, 9}, {15, 7}}, zigzagLevelOrder(root))

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	assert.Equal([][]int{{1}, {3, 2}, {4, 5}}, zigzagLevelOrder(root))
}
