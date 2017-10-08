package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	assert := assert.New(t)

	/*
	     25
	    / \
	   20  35
	   /   /
	  10  30
	*/
	excepted := &TreeNode{Val: 25, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 35, Left: &TreeNode{Val: 30}}}
	head := &ListNode{Val: 10, Next: &ListNode{Val: 20, Next: &ListNode{Val: 25, Next: &ListNode{Val: 30, Next: &ListNode{Val: 35}}}}}
	assert.Equal(excepted, sortedListToBST(head))
}
