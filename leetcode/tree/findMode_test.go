package leetcode

import "testing"

func Test_findMode(t *testing.T) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}
	if r := findMode(root); len(r) != 1 || r[0] != 2 {
		t.Fatal(r)
	}
}
