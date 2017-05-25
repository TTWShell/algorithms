package leetcode

import "testing"

func Test_rob(t *testing.T) {
	input := [][]int{[]int{1, 2, 3}, []int{1, 2, 5, 7, 4, 3, 1, 3, 11, 3, 4, 2, 1, 8, 4, 11, 2, 2, 34}}
	result := []int{4, 80}

	for i := 0; i < len(input); i++ {
		if r := rob(input[i]); r != result[i] {
			t.Error(input[i], result[i], r)
		}
	}
}
