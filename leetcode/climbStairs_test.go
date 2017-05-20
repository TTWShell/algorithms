package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	nums := []int{-1, 0, 1, 2, 3, 4}
	results := []int{0, 0, 1, 2, 3, 5}
	for i := 0; i < len(nums); i++ {
		if result := climbStairs(nums[i]); result != results[i] {
			t.Log(nums[i], results[i], result)
			t.Fail()
		}
	}
}
