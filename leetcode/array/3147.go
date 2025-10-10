package larray

import (
	"slices"
)

func maximumEnergy(energy []int, k int) int {
	length := len(energy)

	dp := make([]int, length)
	for i := length - k; i < length; i++ {
		dp[i] = energy[i]
	}

	for i := length - k - 1; i >= 0; i-- {
		dp[i] = dp[i+k] + energy[i]

	}
	return slices.Max(dp)
}
