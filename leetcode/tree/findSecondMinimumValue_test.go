package leetcode

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}}
	if r := findSecondMinimumValue(root); r != 5 {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	if r := findSecondMinimumValue(root); r != -1 {
		t.Fatal(r)
	}
}
