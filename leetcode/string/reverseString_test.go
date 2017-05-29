package leetcode

import "testing"

func Test_reverseString(t *testing.T) {
	if r := reverseString("hello 中国"); r != "国中 olleh" {
		t.Fatal("hello-->", r)
	}
}
