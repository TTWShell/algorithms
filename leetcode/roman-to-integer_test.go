package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	nums := []string{"DCXXI"}
	result := []int{621}
	for index, num := range nums {
		if romanToInt(num) != result[index] {
			t.Log(num, romanToInt(num), result[index])
			t.Fail()
		}
	}
}
