package lstring

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	input := []string{"a", "b"}
	if result := longestCommonPrefix(input); result != "" {
		t.Fatal(input, result, "")
	}
	input = []string{"ab", "a"}
	if result := longestCommonPrefix(input); result != "a" {
		t.Fatal(input, result, "a")
	}
}
