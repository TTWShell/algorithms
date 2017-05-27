package leetcode

import "testing"

func Test_invertTree(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}

	if r := invertTree(root); r.String() != "<<<<nil> 9 <nil>> 7 <<nil> 6 <nil>>> 4 <<<nil> 3 <nil>> 2 <<nil> 1 <nil>>>>" {
		t.Fatal(r)
	}
}
