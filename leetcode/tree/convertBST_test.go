package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convertBST(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Val: 2}, Val: 5, Right: &TreeNode{Val: 13}}
	excepted := &TreeNode{Left: &TreeNode{Val: 20}, Val: 18, Right: &TreeNode{Val: 13}}
	assert.Equal(excepted, convertBST(root))

	root = &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: -4}, Val: 0, Right: &TreeNode{Val: 1}}, Val: 2, Right: &TreeNode{Val: 3}}
	excepted = &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 2}, Val: 6, Right: &TreeNode{Val: 6}}, Val: 5, Right: &TreeNode{Val: 3}}
	assert.Equal(excepted, convertBST(root))
}
