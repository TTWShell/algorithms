package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3, 2, 1}, postorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
