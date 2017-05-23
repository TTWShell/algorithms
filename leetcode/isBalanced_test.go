package leetcode

import "testing"

func Test_isBalanced(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	if r := isBalanced(root); r != true {
		t.Log(root, r)
		t.Fail()
	}
}
