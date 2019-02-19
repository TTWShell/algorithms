/* https://leetcode.com/problems/powerful-integers/description/
Given two non-negative integers x and y, an integer is powerful if it is equal to x^i + y^j for some integers i >= 0 and j >= 0.

Return a list of all powerful integers that have value less than or equal to bound.

You may return the answer in any order.  In your answer, each value should occur at most once.


Example 1:

	Input: x = 2, y = 3, bound = 10
	Output: [2,3,4,5,7,9,10]
	Explanation:
	2 = 2^0 + 3^0
	3 = 2^1 + 3^0
	4 = 2^0 + 3^1
	5 = 2^1 + 3^1
	7 = 2^2 + 3^1
	9 = 2^3 + 3^0
	10 = 2^0 + 3^2

xample 2:

	Input: x = 3, y = 5, bound = 15
	Output: [2,4,6,8,10,14]

Note:

	1 <= x <= 100
	1 <= y <= 100
	0 <= bound <= 10^6
*/

package lmath

import "fmt"

func powerfulIntegers(x int, y int, bound int) []int {
	cache := map[[2]int]int{[2]int{x, 0}: 1, [2]int{y, 0}: 1}
	power := func(a, b int) int { // return a^b
		if value, ok := cache[[2]int{a, b}]; ok {
			return value
		}
		if value, ok := cache[[2]int{a, b - 1}]; ok {
			res := value * a
			cache[[2]int{a, b}] = res
			return res
		}
		panic(fmt.Sprintf("power err: a = %d, b = %d, %v", a, b, cache))
	}

	resCache := map[int]bool{}
	for i := 0; ; i++ {
		if powerX := power(x, i); (x != 1 || i < 1) && powerX < bound {
			for j := 0; y != 1 || j < 1; j++ {
				if tmp := powerX + power(y, j); tmp <= bound {
					resCache[tmp] = true
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	res, i := make([]int, len(resCache), len(resCache)), 0
	for k := range resCache {
		res[i] = k
		i++
	}
	return res
}
