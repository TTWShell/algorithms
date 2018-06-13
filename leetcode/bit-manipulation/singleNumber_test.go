package lbm

import "testing"

func Test_singleNumber(t *testing.T) {
	input := [][]int{[]int{1}, []int{1, 2, 3, 4, 1, 2, 3}, []int{16, -4, 15, 15, -4}}
	result := []int{1, 4, 16}

	for i := 0; i < len(input); i++ {
		if r := singleNumber(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
