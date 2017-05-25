package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	input := []string{"a", "b"}
	if result := longestCommonPrefix(input); result != "" {
		t.Error(input, result, "")
	}
	input = []string{"ab", "a"}
	if result := longestCommonPrefix(input); result != "a" {
		t.Error(input, result, "a")
	}
}
