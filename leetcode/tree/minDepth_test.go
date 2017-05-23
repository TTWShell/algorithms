package leetcode

import "testing"

func Test_minDepth(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	if r := minDepth(root); r != 2 {
		t.Log(r)
		t.Fail()
	}
}
