package ltree

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}
	if r := trimBST(root, 1, 2); !reflect.DeepEqual(r, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}) {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4}}
	if r := trimBST(root, 1, 3); !reflect.DeepEqual(r, &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}) {
		t.Fatal(r)
	}
}
