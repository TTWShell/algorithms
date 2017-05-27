package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 2, 3, 3, 4, 5, 6}
	dealedNums := []int{1, 2, 3, 4, 5, 6}
	length := removeDuplicates(nums)
	for i := 0; i < len(dealedNums); i++ {
		if length != 6 || nums[i] != dealedNums[i] {
			t.Fatal("return is:", length, nums[i], "!=", dealedNums[i])
		}
	}

}
