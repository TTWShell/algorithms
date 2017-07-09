package leetcode

import "testing"

func Test_averageOfLevels(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if r := averageOfLevels(root); len(r) != 3 || r[0] != 3 || r[1] != 14.5 || r[2] != 11 {
		t.Fatal(r)
	}
}
