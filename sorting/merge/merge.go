package merge

func MergeSort(arr1, arr2 []int) []int {
	len1, len2, lenMerge := len(arr1), len(arr2), len(arr1)+len(arr2)
	mergedArr := make([]int, lenMerge, lenMerge)
	i, j, k := 0, 0, 0

	for ; i < len1 && j < len2; k++ {
		if arr1[i] <= arr2[j] {
			mergedArr[k] = arr1[i]
			i++
		} else {
			mergedArr[k] = arr2[j]
			j++
		}
	}

	// deal tail
	for i < len1 {
		mergedArr[k] = arr1[i]
		k++
		i++
	}
	for j < len2 {
		mergedArr[k] = arr2[j]
		k++
		j++
	}
	return mergedArr
}
