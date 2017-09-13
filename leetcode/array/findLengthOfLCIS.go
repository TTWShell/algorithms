package leetcode

func findLengthOfLCIS(nums []int) int {
	res, tmp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmp++
		} else {
			if tmp > res {
				res = tmp
			}
			tmp = 1
		}
	}
	if len(nums) > 0 && tmp > res {
		return tmp
	}
	return res
}
