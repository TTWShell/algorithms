/* https://leetcode.com/problems/power-of-two/#/description
Given an integer, write a function to determine if it is a power of two.
*/

package leetcode

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

/*
func isPowerOfTwo(n int) bool {
	c := 0
	for n > 0 {
		c += n & 1
		n >>= 1
	}
	return c == 1
}
*/
