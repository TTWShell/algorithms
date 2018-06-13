/* https://leetcode.com/problems/permutation-sequence/description/
The set [1,2,3,…,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order,
We get the following sequence (ie, for n = 3):

    "123"
    "132"
    "213"
    "231"
    "312"
    "321"
Given n and k, return the kth permutation sequence.

Note: Given n will be between 1 and 9 inclusive.
*/

package lbacktracking

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	factorial := func(n int) int {
		sum := 1
		for n > 1 {
			sum *= n
			n--
		}
		return sum
	}

	var helper func(arr []string, k int) string
	helper = func(arr []string, k int) string {
		if k == 1 {
			return strings.Join(arr, "")
		}

		subN := factorial(len(arr) - 1)
		index := k / subN
		if k%subN == 0 {
			index--
		}

		tmpArr := []string{}
		tmpArr = append(tmpArr, arr[:index]...)
		if tmp := index + 1; tmp < len(arr) {
			tmpArr = append(tmpArr, arr[tmp:]...)
		}
		return arr[index] + helper(tmpArr, k-subN*index)
	}

	arr := make([]string, n, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = strconv.Itoa(i)
	}
	return helper(arr, k)
}
