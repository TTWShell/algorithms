package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	assert.Equal([]int{1, 3, 2}, inorderTraversal(root))
}
