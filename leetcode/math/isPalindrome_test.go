package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	nums := []int{-121, 121, 123, 123454321}
	results := []bool{false, true, false, true}
	for index, num := range nums {
		if r := isPalindrome(num); r != results[index] {
			t.Fatal(num, r, results[index])
		}
	}
}
