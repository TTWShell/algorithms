package leetcode

import "testing"

func Test_levelOrderBottom(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	result := [][]int{[]int{15, 7}, []int{9, 20}, []int{3}}

	for i, r := range levelOrderBottom(root) {
		for j := 0; j < len(r); j++ {
			if r[j] != result[i][j] {
				t.Log(levelOrderBottom(root))
				t.Fail()
			}
		}
	}
}
