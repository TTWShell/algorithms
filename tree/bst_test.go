package tree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
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
		t.Fatal("inorder check failed", r)
	}
}

func TestSearch(t *testing.T) {
	bst := NewBST()
	if r := bst.Search(1); r {
		t.Fatal(r, "empty tree should return false")
	}

	els := []int{3, 2, 15, 18, 23, 5, 7}
	for _, el := range els {
		bst.Insert(el)
	}

	for _, el := range els {
		if r := bst.Search(el); !r {
			t.Fatal(r, el)
		}
	}

	if r := bst.Search(143414); r {
		t.Fatal(r, "143414 should be not exist")
	}
}
