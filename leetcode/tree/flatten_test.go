package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flatten(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Val:   1,
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	excepted := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}}

	flatten(root)
	assert.Equal(excepted, root)
}
