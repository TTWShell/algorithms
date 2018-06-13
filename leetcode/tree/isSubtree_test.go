package ltree

import "testing"

func Test_isSubtree(t *testing.T) {
	s := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}
	subt := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	if r := isSubtree(s, subt); r != true {
		t.Fatal(s, subt, r)
	}

	s = &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 5}}
	if r := isSubtree(s, subt); r != false {
		t.Fatal(s, subt, r)
	}
}
