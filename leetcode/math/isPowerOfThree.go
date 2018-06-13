/* https://leetcode.com/problems/power-of-three/#/description
Given an integer, write a function to determine if it is a power of three.

Follow up:
Could you do it without using any loop / recursion?
*/

package lmath

func isPowerOfThree(n int) bool {
	// 0-2^31: 3^19 = 1162261467
	return n > 0 && 1162261467%n == 0
}
