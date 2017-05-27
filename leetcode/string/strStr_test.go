package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	haystacks := []string{"a", "", "abc", "abcdabcdf"}
	needles := []string{"b", "", "b", "abcdf"}
	results := []int{-1, 0, 1, 4}
	for i := 0; i < len(haystacks); i++ {
		if r := strStr(haystacks[i], needles[i]); r != results[i] {
			t.Fatal(haystacks[i], needles[i], results[i], "output is:", r)
		}
	}
}
