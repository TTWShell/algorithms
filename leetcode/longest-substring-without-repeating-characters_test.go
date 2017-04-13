package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	str := []string{"a", "au", "abcabcbb", "bbbbb", "pwwkew"}
	result := []int{1, 2, 3, 1, 3}
	for index, s := range str {
		if lengthOfLongestSubstring(s) != result[index] {
			t.Log(result)
			t.Fail()
		}
	}
}
