package lmath

import "testing"

func Test_countPrimes(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := []int{0, 0, 1, 2, 2, 3, 3, 4, 4, 4}

	for i := 0; i < len(nums); i++ {
		if r := countPrimes(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
