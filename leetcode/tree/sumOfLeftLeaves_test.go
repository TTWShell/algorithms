package leetcode

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if r := sumOfLeftLeaves(root); r != 24 {
		t.Fatal(root, r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	if r := sumOfLeftLeaves(root); r != 4 {
		t.Fatal(root, r)
	}
}
