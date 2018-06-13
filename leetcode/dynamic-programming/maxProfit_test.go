package ldp

import "testing"

func Test_maxProfit(t *testing.T) {
	input := [][]int{[]int{7, 1, 5, 3, 6, 4}, []int{7, 6, 4, 3, 1}}
	result := []int{5, 0}

	for i := 0; i < len(input); i++ {
		if r := maxProfit(input[i]); r != result[i] {
			t.Fatal(input[i], result[i])
		}
	}
}
