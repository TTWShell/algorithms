package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	nums := []int{1, 2, 123, 321, 1234, -3453, 1534236469, -2147483648}
	result := []int{1, 2, 321, 123, 4321, -3543, 0, 0}
	for index, num := range nums {
		if reverse(num) != result[index] {
			t.Log(num, reverse(num), result[index])
			t.Fail()
		}
	}
}
