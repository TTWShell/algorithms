package leetcode

import "testing"

func Test_reverseStr(t *testing.T) {
	if r := reverseStr("abcdefg", 2); r != "bacdfeg" {
		t.Fatal(r)
	}

	if r := reverseStr("abcd", 4); r != "dcba" {
		t.Fatal(r)
	}

	if r := reverseStr("abcdefg", 8); r != "gfedcba" {
		t.Fatal(r)
	}
}
