package leetcode

import "testing"

func Test_isSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	if r := isSymmetric(root); r != true {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	if r := isSymmetric(root); r != false {
		t.Fatal(r)
	}
}
