package leetcode

import "testing"

func Test_tree2str(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	if r := tree2str(tree); r != "1(2(4))(3)" {
		t.Fatal(r)
	}

	tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	if r := tree2str(tree); r != "1(2()(4))(3)" {
		t.Fatal(r)
	}
}
