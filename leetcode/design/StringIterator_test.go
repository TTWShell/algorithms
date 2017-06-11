package leetcode

import "testing"

func Test_StringIterator(t *testing.T) {
	iterator := Constructor("L1e2")
	if next := iterator.Next(); next != 'L' {
		t.Fatal(string(next))
	}

	if HasNext := iterator.HasNext(); HasNext != true {
		t.Fatal("HasNext error.")
	}

	iterator.Next()
	iterator.Next()
	if HasNext := iterator.HasNext(); HasNext != false {
		t.Fatal(HasNext)
	}
	if next := iterator.Next(); next != ' ' {
		t.Fatal("isempty must return a white space.")
	}
}
