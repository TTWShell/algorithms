package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	ss := []string{"", " ", "a ", "b a ", "hello world"}
	results := []int{0, 0, 1, 1, 5}
	for i := 0; i < len(ss); i++ {
		if r := lengthOfLastWord(ss[i]); r != results[i] {
			t.Error(ss[i], results[i], r)
		}
	}
}
