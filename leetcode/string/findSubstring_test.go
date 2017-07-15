package leetcode

import "testing"

func Test_findSubstring(t *testing.T) {
	if r := findSubstring("barfoothefoobarman", []string{"foo", "bar"}); len(r) != 2 ||
		r[0] != 0 || r[1] != 9 {
		t.Fatal(r)
	}
}
