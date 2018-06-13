/* https://leetcode.com/problems/nth-digit/#/description
Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...

Note:
n is positive and will fit within the range of a 32-bit signed integer (n < 2^31).

Example 1:

    Input:
    3

    Output:
    3
Example 2:

    Input:
    11

    Output:
    0

Explanation:
The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
1-9			9
10-99		90 * 2
100-999		900 * 3
1000-9999	9000 * 4
*/

package lmath

func findNthDigit(n int) int {
	// 找出第n个digit所在的数
	temp, digit := 0, 1
	for ; temp < n; digit++ {
		temp += 9 * powForfindNthDigit(10, digit-1) * digit
	}
	digit -= 1
	temp -= 9 * powForfindNthDigit(10, digit-1) * digit
	num := powForfindNthDigit(10, digit-1) - 1 + (n-temp)/digit // 待修正

	// 找到第n个digit在num中的index
	// 如果index=0，则表示目标digit在num的最后位置，否则在对应的index位置，所以实际index ＝index －1
	index := (n - temp) % digit
	if index != 0 {
		num += 1
	} else {
		index = digit
	}

	// 找到目标digit，按下标取之
	for i := index; i < digit; i++ {
		num /= 10
	}
	return num % 10
}

func powForfindNthDigit(x, y int) int {
	r := 1
	for i := 0; i < y; i++ {
		r *= x
	}
	return r
}
