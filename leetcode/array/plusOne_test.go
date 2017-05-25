package leetcode

import "testing"

func Test_plusOne(t *testing.T) {
	digits := []int{9, 9, 9}
	result := []int{1, 0, 0, 0}
	r := plusOne(digits)
	for i := 0; i < len(r); i++ {
		if r[i] != result[i] {
			t.Error(result, r)
		}
	}
}
