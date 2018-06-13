package lstring

import "testing"

func Test_firstUniqChar(t *testing.T) {
	input := []string{"leetcode", "loveleetcode", "aabbcc"}
	result := []int{0, 2, -1}

	for i := range input {
		if r := firstUniqChar(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
