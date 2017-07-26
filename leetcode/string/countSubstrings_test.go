package leetcode

import "testing"

func Test_countSubstrings(t *testing.T) {
	if r := countSubstrings("abc"); r != 3 {
		t.Fatal(r)
	}

	if r := countSubstrings("aaa"); r != 6 {
		t.Fatal(r)
	}
}
