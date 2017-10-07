package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal([][]int{
		{3},
		{9, 20},
		{15, 7},
	}, levelOrder(root))
}
