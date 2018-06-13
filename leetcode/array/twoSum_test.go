package larray

import "testing"

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	if result[0] != 0 || result[1] != 1 {
		t.Fatal(result)
	}
}
