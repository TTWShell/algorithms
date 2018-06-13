/* https://leetcode.com/problems/number-complement/#/description
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.

Note:
    The given integer is guaranteed to fit within the range of a 32-bit signed integer.
    You could assume no leading zero bit in the integer’s binary representation.
Example 1:
    Input: 5
    Output: 2
    Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.

Example 2:
    Input: 1
    Output: 0
    Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
*/

package lbm

func findComplement(num int) int {
	var mask int64 = 1
	for v := num; v > 0; v = v >> 1 {
		mask = mask << 1
	}
	num = ^num & int(mask-1)
	return num
}

/*
func findComplement(num int) int {
	bits := make([]int, 32, 32)
	for i := 31; num > 0; i-- {
		bits[i] = num % 2
		num = num >> 1
	}

	pow := func(x, y int) int {
		r := 1
		for i := 1; i < y; i++ {
			r *= x
		}
		return r
	}

	tag, r := false, 0
	for i := 0; i < 32; i++ {
		if tag == false && bits[i] == 1 {
			tag = true
		}
		if tag == true && bits[i] == 0 {
			r += pow(2, 32-i)
		}
	}
	return r
}
*/
