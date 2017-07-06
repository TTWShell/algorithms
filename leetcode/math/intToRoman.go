/* https://leetcode.com/problems/integer-to-roman/#/description
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
*/

package leetcode

func intToRoman(num int) string {
	var (
		M = []string{"", "M", "MM", "MMM"}
		C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	)
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}
