package leetcode

import "testing"

func Test_toHex(t *testing.T) {
	nums := []int{0, 16, 26, -1, -3}
	results := []string{"0", "10", "1a", "ffffffff", "fffffffd"}

	for i := range nums {
		if r := toHex(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
