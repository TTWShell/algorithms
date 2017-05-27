package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	input := []string{"A man, a plan, a canal: Panama", "race a car"}
	result := []bool{true, false}

	for i := 0; i < len(input); i++ {
		if r := isPalindrome(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
