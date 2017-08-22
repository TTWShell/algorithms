package shell

func ShellSort(arr []int) {
	for k := len(arr) / 2; k >= 1; k = k / 2 {
		for i := k; i < len(arr); i++ {
			tmp := arr[i]
			j := i - k
			for j >= 0 && arr[j] > tmp {
				arr[j+k] = arr[j]
				j -= k
			}
			arr[j+k] = tmp
		}
	}
}
