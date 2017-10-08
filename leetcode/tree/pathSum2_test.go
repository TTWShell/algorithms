package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum2(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Val:   5,
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
	}

	assert.Equal([][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum2(root, 22))
}
