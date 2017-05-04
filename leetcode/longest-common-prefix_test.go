package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	input := []string{"a", "b"}
	if longestCommonPrefix(input) != "" {
		t.Log(input, longestCommonPrefix(input), "")
		t.Fail()
	}
	input = []string{"ab", "a"}
	if longestCommonPrefix(input) != "a" {
		t.Log(input, longestCommonPrefix(input), "a")
		t.Fail()
	}
}
