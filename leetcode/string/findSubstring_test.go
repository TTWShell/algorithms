package leetcode

import "testing"

func Test_findSubstring(t *testing.T) {
	if r := findSubstring("abaababbaba", []string{"ab", "ba", "ab", "ba"}); len(r) != 2 || r[0] != 1 || r[1] != 3 {
		t.Fatal(r)
	}

	if r := findSubstring("aaa", []string{"a", "a"}); len(r) != 2 || r[0] != 0 || r[1] != 1 {
		t.Fatal(r)
	}

	if r := findSubstring("barfoothefoobarman", []string{"foo", "bar"}); len(r) != 2 ||
		r[0] != 0 || r[1] != 9 {
		t.Fatal(r)
	}

	if r := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}); len(r) != 1 || r[0] != 8 {
		t.Fatal(r)
	}

	if r := findSubstring("a", []string{"a"}); len(r) != 1 || r[0] != 0 {
		t.Fatal(r)
	}
}
