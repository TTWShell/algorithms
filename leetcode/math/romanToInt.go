/* https://leetcode.com/problems/roman-to-integer/#/description
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

http://baike.baidu.com/item/罗马数字
*/
package leetcode

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	cur, last := 0, roman[s[0]]
	sum := last
	for i := 1; i < len(s); i++ {
		cur = roman[s[i]]
		if cur > last {
			sum += cur - 2*last
		} else {
			sum += cur
		}
		last = cur
	}
	return sum
}
