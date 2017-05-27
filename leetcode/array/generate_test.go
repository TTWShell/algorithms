package leetcode

import "testing"

func Test_generate(t *testing.T) {
	result := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}
	r := generate(5)

	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] != result[i][j] {
				t.Fatal(r, result)
			}
		}
	}
}
