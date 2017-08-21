package quick

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	guard := 0
	for i, tag := 1, arr[0]; i < len(arr); i++ {
		if arr[i] < tag {
			arr[i], arr[guard+1] = arr[guard+1], arr[i]
			arr[guard], arr[guard+1] = arr[guard+1], arr[guard]
			guard++
		}
	}

	if guard > 0 {
		QuickSort(arr[:guard])
	}
	if guard < len(arr) {
		QuickSort(arr[guard+1:])
	}
}
