package leetcode

import "testing"

func Test_toHex(t *testing.T) {
	nums := []int{26, -1}
	results := []string{"1a", "ffffffff"}

	for i := range nums {
		if r := toHex(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
