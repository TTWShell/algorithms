package search

func BinarySearch(sortedData []int, target int) (index int) {
	start, end := 0, len(sortedData)-1

	if target < sortedData[start] || target > sortedData[end] {
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2
		switch {
		case target < sortedData[mid]:
			end = mid - 1
		case target > sortedData[mid]:
			start = mid + 1
		case target == sortedData[mid]:
			return mid
		}
	}
	return -1
}
