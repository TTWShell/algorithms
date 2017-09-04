package tree

import (
	"reflect"
	"testing"
)

func TestInsertBST(t *testing.T) {
	bst := NewBST()
	if r := bst.Insert(1); !r || !reflect.DeepEqual(bst.root, &TreeNode{Val: 1}) {
		t.Fatal(bst.root, r)
	}

	els := []int{3, 2, 15, 18, 23, 5, 7}
	for _, el := range els {
		if r := bst.Insert(el); !r {
			t.Fatal(bst.root, r)
		}
	}

	if r := bst.Insert(1); r {
		t.Fatal(bst.root, r)
	}

	if r := ReOutput(DFSInOrder, bst.root); r != "1 2 3 5 7 15 18 23" {
		t.Fatal("中序遍历检查失败", r)
	}
}
