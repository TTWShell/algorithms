package lht

import "testing"

func Test_containsDuplicate(t *testing.T) {
	input := [][]int{[]int{1, 2, 3}, []int{1, 2, 2}}
	result := []bool{false, true}

	for i, nums := range input {
		if r := containsDuplicate(nums); r != result[i] {
			t.Fatal(nums, r)
		}
	}
}
