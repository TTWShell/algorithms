package leetcode

import "testing"

func Test_majorityElement(t *testing.T) {
	nums := []int{1, 2, 3, 4, 2, 2, 2}
	if r := majorityElement(nums); r != 2 {
		t.Fatal(nums, r)
	}
	nums = []int{1}
	if r := majorityElement(nums); r != 1 {
		t.Fatal(nums, r)
	}
}
