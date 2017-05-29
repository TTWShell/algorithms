package leetcode

import "testing"

func Test_reverseString(t *testing.T) {
	if r := reverseString("hello"); r != "olleh" {
		t.Fatal("hello-->", r)
	}
}
