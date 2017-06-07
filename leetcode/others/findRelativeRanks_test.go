package leetcode

import "testing"

func Test_findRelativeRanks(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	output := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}

	r := findRelativeRanks(input)
	for i := range output {
		if r[i] != output[i] {
			t.Fatal(r)
		}
	}
}
