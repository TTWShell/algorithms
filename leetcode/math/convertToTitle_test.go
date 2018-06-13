package lmath

import "testing"

func Test_convertToTitle(t *testing.T) {
	input := []int{1, 2, 3, 26, 27, 28}
	result := []string{"A", "B", "C", "Z", "AA", "AB"}

	for i := 0; i < len(input); i++ {
		if r := convertToTitle(input[i]); r != result[i] {
			t.Fatal(input[i], result[i], r)
		}
	}
}
