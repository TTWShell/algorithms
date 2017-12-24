package insertion

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > value; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = value
	}
}
