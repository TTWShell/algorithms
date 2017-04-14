package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	nums := []int{-121, 121, 123, 123454321}
	result := []bool{false, true, false, true}
	for index, num := range nums {
		if isPalindrome(num) != result[index] {
			t.Log(num, isPalindrome(num), result[index])
			t.Fail()
		}
	}
}
