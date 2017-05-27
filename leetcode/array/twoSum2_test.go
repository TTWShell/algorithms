package leetcode

import "testing"

func Test_twoSum2(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	if r := twoSum2(numbers, target); r[0] != 1 || r[1] != 2 {
		t.Fatal(numbers, target, r)
	}
}
