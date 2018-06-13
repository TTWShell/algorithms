package ltree

import "testing"

func Test_mergeTrees(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}

	result := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}}

	if r := isSameTree(result, mergeTrees(t1, t2)); r != true {
		t.Fatal(r)
	}
}
