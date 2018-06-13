package ltree

import "testing"

func Test_isSameTree(t *testing.T) {
	p := &TreeNode{Val: 1}
	q := &TreeNode{Val: 1}
	if r := isSameTree(p, q); r != true {
		t.Fatal(p, r, q)
	}

	p = &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}
	q = &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 1}}
	if r := isSameTree(p, q); r != false {
		t.Fatal(p, r, q)
	}
}
