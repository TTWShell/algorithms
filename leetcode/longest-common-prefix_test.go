package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	input := []string{"a", "b"}
	if result := longestCommonPrefix(input); result != "" {
		t.Log(input, result, "")
		t.Fail()
	}
	input = []string{"ab", "a"}
	if result := longestCommonPrefix(input); result != "a" {
		t.Log(input, result, "a")
		t.Fail()
	}
}
