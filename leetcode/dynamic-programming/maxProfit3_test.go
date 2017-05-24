package leetcode

import "testing"

func Test_maxProfit3(t *testing.T) {
	input := [][]int{[]int{7, 1, 5, 3, 6, 4}, []int{1, 2, 1, 2, 4, 2, 5, 3}}
	result := []int{7, 6}

	for i := 0; i < len(input); i++ {
		if r := maxProfit3(input[i]); r != result[i] {
			t.Log(input[i], r, "Expected:", result[i])
			t.Fail()
		}
	}
}
