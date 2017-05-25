package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	if r := sortedArrayToBST(nums); r.String() != "<<<<nil> 1 <nil>> 2 <<nil> 3 <nil>>> 4 <<<nil> 5 <nil>> 6 <<nil> 7 <nil>>>>" {
		t.Error(nums, r.String())
	}
}
