package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	assert := assert.New(t)

	/*
	     3
	    / \
	   9  20
	     /  \
	    15   7
	*/
	excepted := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	assert.Equal(excepted, buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
