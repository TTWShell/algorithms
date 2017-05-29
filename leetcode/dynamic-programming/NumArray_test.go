package leetcode

import "testing"

func Test_NumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}

	numarray := Constructor(nums)

	if r := numarray.SumRange(0, 2); r != 1 {
		t.Fatal("SumRange(0, 2)", r)
	}

	if r := numarray.SumRange(2, 5); r != -1 {
		t.Fatal("SumRange(2, 5)", r)
	}

	if r := numarray.SumRange(0, 5); r != -3 {
		t.Fatal("SumRange(0, 5)", r)
	}
}
