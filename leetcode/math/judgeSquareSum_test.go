package leetcode

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	input := []int{1, 2, 3, 5}
	output := []bool{true, true, false, true}

	for i := range input {
		if r := judgeSquareSum(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
