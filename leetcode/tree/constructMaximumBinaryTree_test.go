package leetcode

import "testing"

func Test_constructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	output := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 0},
		},
	}
	if r := constructMaximumBinaryTree(nums); r.String() != output.String() {
		t.Fatal(r)
	}
}
