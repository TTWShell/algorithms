package leetcode

import "testing"

func Test_reverseString(t *testing.T) {
	if r := reverseString("hello china中国"); r != "国中anihc olleh" {
		t.Fatal("hello-->", r)
	}
}
