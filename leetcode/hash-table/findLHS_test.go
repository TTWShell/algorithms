package lht

import "testing"

func Test_findLHS(t *testing.T) {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	if r := findLHS(nums); r != 5 {
		t.Fatal(nums, r)
	}
}
