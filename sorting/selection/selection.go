package selection

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		max, index := arr[i], i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < max {
				max, index = arr[j], j
			}
		}
		if i != index {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}
