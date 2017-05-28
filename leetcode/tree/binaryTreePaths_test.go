package leetcode

import "testing"

func Test_binaryTreePaths(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	result := []string{"1->2->5", "1->3"}

	r := binaryTreePaths(root)

	for i, s := range result {
		if s != r[i] {
			t.Fatal(r)
		}
	}
}
