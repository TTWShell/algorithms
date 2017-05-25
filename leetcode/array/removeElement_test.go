package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	input := []int{3, 2, 2, 4, 3, 1, 4, 5, 5, 6}
	result := []int{2, 2, 4, 1, 4, 5, 5, 6}
	length := removeElement(input, 3)
	for i := 0; i < len(result); i++ {
		if length != 8 || input[i] != result[i] {
			t.Error(length, input[i], result[i])
		}
	}
}
