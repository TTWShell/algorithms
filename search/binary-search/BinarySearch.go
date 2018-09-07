package search

func search(nums []int, target int) (index int) {
	start, end := 0, len(nums)-1

	if target < nums[start] || target > nums[end] {
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2
		switch {
		case target < nums[mid]:
			end = mid - 1
		case target > nums[mid]:
			start = mid + 1
		case target == nums[mid]:
			return mid
		}
	}
	return -1
}
