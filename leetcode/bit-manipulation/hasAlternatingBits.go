/* https://leetcode.com/problems/binary-number-with-alternating-bits/description/
Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example 1:
    Input: 5
    Output: True
    Explanation:
    The binary representation of 5 is: 101
Example 2:
    Input: 7
    Output: False
    Explanation:
    The binary representation of 7 is: 111.
Example 3:
    Input: 11
    Output: False
    Explanation:
    The binary representation of 11 is: 1011.
Example 4:
    Input: 10
    Output: True
    Explanation:
    The binary representation of 10 is: 1010.
*/
package lbm

// others: https://discuss.leetcode.com/category/1550/binary-number-with-alternating-bits
func hasAlternatingBits(n int) bool {
	if n ^= n / 4; n&(n-1) > 0 {
		return false
	}
	return true
}
