package leetcode

import (
	"reflect"
	"testing"
)

func Test_printTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	output := [][]string{{"", "1", ""}, {"2", "", ""}}
	if r := printTree(root); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	output = [][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}}
	if r := printTree(root); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 5}}
	output = [][]string{
		{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
		{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
		{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	}
	if r := printTree(root); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}
}
