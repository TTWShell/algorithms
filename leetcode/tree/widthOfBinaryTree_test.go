package leetcode

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}}
	if r := widthOfBinaryTree(root); r != 4 {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}}
	if r := widthOfBinaryTree(root); r != 2 {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	if r := widthOfBinaryTree(root); r != 2 {
		t.Fatal(r)
	}

	root = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 7}}},
	}
	if r := widthOfBinaryTree(root); r != 8 {
		t.Fatal(r)
	}

	if r := widthOfBinaryTree(nil); r != 0 {
		t.Fatal(r)
	}

	if r := widthOfBinaryTree(&TreeNode{Val: 1}); r != 1 {
		t.Fatal(r)
	}
}
