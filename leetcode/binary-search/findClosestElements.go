/* https://leetcode.com/problems/find-k-closest-elements/description/
Given a sorted array, two integers k and x, find the k closest elements to x in the array. The result should also be sorted in ascending order.
If there is a tie, the smaller elements are always preferred.

Example 1:
	Input: [1,2,3,4,5], k=4, x=3
	Output: [1,2,3,4]

Example 2:
	Input: [1,2,3,4,5], k=4, x=-1
	Output: [1,2,3,4]

Note:
	1. The value k is positive and will always be smaller than the length of the sorted array.
	2. Length of the given array is positive and will not exceed 10^4
	3. Absolute value of elements in the array and x will not exceed 10^4
*/

package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	binarySearch := func(arr []int, target int) int {
		// return -1: less min; -2: more max; 0~len(arr)-1: index if exist else first num's index(num < target)
		if target < arr[0] {
			return -1
		}
		if target > arr[len(arr)-1] {
			return -2
		}

		start, end, mid := 0, len(arr)-1, 0
		for start <= end {
			mid = start + (end-start)/2
			if arr[mid] == target {
				return mid
			} else if arr[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return mid
	}

	index := binarySearch(arr, x)

	if index == -1 || index == 0 {
		return arr[0:k]
	}
	if index == -2 {
		return arr[len(arr)-k:]
	}

	start, end := 0, len(arr)-1
	if index-k > start {
		start = index - k
	}
	if index+k < end {
		end = index + k
	}

	for end-start >= k {
		if x-arr[start] <= arr[end]-x {
			end--
		} else if x-arr[start] > arr[end]-x {
			start++
		}
	}

	return arr[start : end+1]
}
