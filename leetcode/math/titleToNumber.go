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

import "math"

func titleToNumber(s string) int {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maps := map[byte]int{}
	for i := 0; i < len(chars); i++ {
		maps[chars[i]] = i + 1
	}
	n := len(s)
	var r int
	for i := 0; i < n; i++ {
		if num, ok := maps[s[i]]; ok {
			r += num * int(math.Pow(float64(26), float64(n-1-i)))
		}
	}
	return r
}
