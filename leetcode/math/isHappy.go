/* https://leetcode.com/problems/happy-number/#/description
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.

Example: 19 is a happy number

12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

对于某一个正整数，如果对其各个位上的数字分别平方，然后再加起来得到一个新的数字，
再进行同样的操作，如果最终结果变成了1，则说明是快乐数。
如果一直循环但不是1的话，就不是快乐数。
*/

package leetcode

// 数学解法：https://en.wikipedia.org/wiki/Happy_number
func isHappy(n int) bool {
	for n != 1 && n != 4 {
		num := 0
		for n > 0 {
			mod := n % 10
			num += mod * mod
			n /= 10
		}
		n = num
	}
	return n == 1
}

/*
func isHappy(n int) bool {
	maps := make(map[int]int)
	for n != 1 {
		r := 0
		for n > 0 {
			mod := n % 10
			r += mod * mod
			n /= 10
		}
		n = r
		if _, ok := maps[n]; ok {
			break
		} else {
			maps[n] = 1
		}
	}
	return n == 1
}
*/
