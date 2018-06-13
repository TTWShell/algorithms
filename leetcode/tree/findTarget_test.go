package ltree

import "testing"

func Test_findTarget(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}

	if r := findTarget(root, 9); !r {
		t.Fatal(r)
	}

	if r := findTarget(root, 28); r {
		t.Fatal(r)
	}
}
