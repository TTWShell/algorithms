package leetcode

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	input := []int{3, 5}
	output := []bool{false, true}

	for i := range input {
		if r := judgeSquareSum(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
