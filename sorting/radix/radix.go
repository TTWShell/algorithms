package radix

func maxBit(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	bit := 1
	for max >= 10 {
		max /= 10
		bit++
	}
	return bit
}

func RadixSort(arr []int) {
	mBit := maxBit(arr)
	tmp := make([]int, len(arr), len(arr))
	count := make([]int, 10, 10) // 计数器
	for i, radix := 1, 1; i <= mBit; i++ {
		// 每次分配前清空计数器
		for j := 0; j < 10; j++ {
			count[j] = 0
		}

		for j := 0; j < len(arr); j++ {
			k := (arr[j] / radix) % 10 // 统计每个桶中的记录数
			count[k]++
		}

		for j := 1; j < 10; j++ {
			count[j] = count[j-1] + count[j] // 将tmp中的位置依次分配给每个桶
		}

		for j := len(arr) - 1; j >= 0; j-- {
			// 将所有桶中记录依次收集到tmp中
			k := (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}

		copy(arr, tmp)
		radix *= 10
	}
}
