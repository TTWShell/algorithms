/* https://leetcode.com/problems/plus-one/#/description
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.
给定一个非负整数，表示为非空数组，然后执行+1运算。
*/

package leetcode

func plusOne(digits []int) []int {
	carry, length := 1, len(digits) // carry default = 1 for units + 1

	for i := length - 1; i >= 0; i-- {
		cur := digits[i] + carry
		if cur == 10 {
			carry = 1
			digits[i] = 0
		} else {
			digits[i] = cur
			return digits
		}
	}

	if carry == 1 {
		r := make([]int, length+1)
		r[0] = 1
		for i := 1; i < length+1; i++ {
			r[i] = digits[i-1]
		}
		return r
	}
	return digits
}
