package leetcode

import "testing"

func Test_maxProfit2(t *testing.T) {
	input := [][]int{[]int{7, 1, 5, 3, 6, 4}, []int{1, 2}}
	result := []int{7, 1}

	for i := 0; i < len(input); i++ {
		if r := maxProfit2(input[i]); r != result[i] {
			t.Error(input[i], r, "Expected:", result[i])
		}
	}
}
