package leetcode

import "testing"

func Test_isHappy(t *testing.T) {
	nums := []int{19, 11}
	result := []bool{true, false}

	for i := 0; i < len(nums); i++ {
		if r := isHappy(nums[i]); r != result[i] {
			t.Error(nums[i], result[i], r)
		}
	}
}
