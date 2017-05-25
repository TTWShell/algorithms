package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 8, Right: &TreeNode{Val: 4}}}
	sums := []int{1, 9, 17, 22}
	results := []bool{false, true, true, false}

	for i := 0; i < len(sums); i++ {
		if r := hasPathSum(root, sums[i]); r != results[i] {
			t.Error(sums[i], r)
		}
	}
}
