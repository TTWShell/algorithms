package leetcode

import "testing"

func Test_findTheDifference(t *testing.T) {
	if r := findTheDifference("abcd", "abcde"); r != 'e' {
		t.Fatal(r)
	}
}
