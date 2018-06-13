package ltree

import "testing"

func Test_findTilt(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if r := findTilt(root); r != 1 {
		t.Fatal(root, r)
	}
}
