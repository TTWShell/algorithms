package leetcode

import "testing"

func Test_getRow(t *testing.T) {
	result := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}
	for i := 0; i < len(result); i++ {
		r := getRow(i)
		for j := 0; j < len(r); j++ {
			if r[j] != result[i][j] {
				t.Error(i-1, result[i], r)
			}
		}
	}
}
