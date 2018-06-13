package lmath

import "testing"

func Test_addDigits(t *testing.T) {
	nums := []int{38, 22, 138}
	results := []int{2, 4, 3}

	for i := range nums {
		if r := addDigits(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
