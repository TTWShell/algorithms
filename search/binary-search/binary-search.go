package main

import "fmt"

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

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, target := range data {
		result := BinarySearch(data, target)
		if index != result {
			fmt.Printf("error: target is number `%d`, index is %d, result is `%d`.\n", target, index, target)
		}
	}
	if BinarySearch(data, 10) != -1 || BinarySearch(data, 0) != -1 {
		fmt.Printf("error when search target not in sortedData")
	}
}
