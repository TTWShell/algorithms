package tree

import "testing"

func TestAVLInsert(t *testing.T) {
	avl := NewAVL()

	els := []int{3, 2, 15, 18, 23, 5, 7}
	for _, el := range els {
		if r := avl.Insert(el); !r {
			t.Fatal(avl.root, r)
		}
		t.Log(avl.root)
	}
}
