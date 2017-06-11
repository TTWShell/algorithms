/* https://leetcode.com/problems/can-place-flowers/#/description
Suppose you have a long flowerbed in which some of the plots are planted and some are not.
However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.

Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty),
and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.

Example 1:
    Input: flowerbed = [1,0,0,0,1], n = 1
    Output: True

Example 2:
    Input: flowerbed = [1,0,0,0,1], n = 2
    Output: False

Note:
    The input array won't violate no-adjacent-flowers rule.
    The input array size is in the range of [1, 20000].
    n is a non-negative integer which won't exceed the input array size.
*/

package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	length, prev, next, count := len(flowerbed), 0, 0, 0

	for i := 0; i < length && count < n; i++ {
		if i == 0 {
			prev = 0
		} else {
			prev = flowerbed[i-1]
		}

		if i == length-1 {
			next = 0
		} else {
			next = flowerbed[i+1]
		}

		if prev == 0 && flowerbed[i] == 0 && next == 0 {
			flowerbed[i] = 1
			count++
		}
	}

	return count == n
}
