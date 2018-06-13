package ltree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	if r := diameterOfBinaryTree(root); r != 3 {
		t.Fatal(r)
	}
}
