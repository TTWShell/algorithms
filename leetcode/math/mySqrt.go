/* https://leetcode.com/problems/sqrtx/#/description
Implement int sqrt(int x).

Compute and return the square root of x.
*/

package lmath

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	start, end := 0, x/2
	for start <= end {
		mid := start + (end-start)/2
		square := mid * mid
		if square < x {
			start = mid + 1
		} else if square == x {
			return mid
		} else {
			end = mid - 1
		}
	}
	return end
}
