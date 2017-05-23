package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	results := []int{2, 3}
	trees := []*TreeNode{tree1, tree2}

	for i, tree := range trees {
		if r := maxDepth(tree); r != results[i] {
			t.Log(tree, results[i])
			t.Fail()
		}
	}
}
