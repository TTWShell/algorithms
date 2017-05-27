package leetcode

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	input := [][]int{{-1, 0, 1, -1}, {1, 2, 3, 1, 3}}
	result := []bool{true, true}

	for i := range input {
		if r := containsNearbyDuplicate(input[i], 3); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
