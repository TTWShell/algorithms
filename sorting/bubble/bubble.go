package bubble

func BubbleSort(arr []int) {
	for stepCount := len(arr) - 1; ; stepCount-- {
		swap := false
		for i := 1; i <= stepCount; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
