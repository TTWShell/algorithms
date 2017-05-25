/* https://leetcode.com/problems/excel-sheet-column-number/#/description
Related to question Excel Sheet Column Title

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
*/

package leetcode

func titleToNumber(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i] - 'A' + 1)
		r += num * pow(26, len(s)-i-1)
	}
	return r
}

func pow(x, y int) int {
	r := 1
	for i := 0; i < y; i++ {
		r *= x
	}
	return r
}
