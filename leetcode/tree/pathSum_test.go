package leetcode

import "testing"

func Test_pathSum(t *testing.T) {
	root := &TreeNode{Val: 10, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}}

	if r := pathSum(root, 8); r != 3 {
		t.Fatal(r, "!=", 3)
	}
}
