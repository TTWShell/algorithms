package ltree

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}
	if r := getMinimumDifference(root); r != 1 {
		t.Fatal(root, r)
	}

	// 236,104,701,null,227,null,911
	root = &TreeNode{Val: 236, Left: &TreeNode{Val: 104, Right: &TreeNode{Val: 227}}, Right: &TreeNode{Val: 701, Right: &TreeNode{Val: 911}}}
	if r := getMinimumDifference(root); r != 9 {
		t.Fatal(root, r)
	}
}
