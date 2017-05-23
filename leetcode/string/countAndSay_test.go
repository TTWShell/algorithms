package leetcode

import "testing"

func Test_countAndSay(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	results := []string{"1", "11", "21", "1211", "111221"}
	for i := 1; i < len(nums); i++ {
		if r := countAndSay(nums[i]); r != results[i] {
			t.Log(nums[i], results[i], "output is:", r)
			t.Fail()
		}
	}
}
