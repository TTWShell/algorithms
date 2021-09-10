/* https://leetcode.com/problems/excel-sheet-column-title/#/description
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
*/

package lmath

func convertToTitle(n int) string {
	r := ""
	for n > 0 {
		r = string(rune((n-1)%26+'A')) + r
		n = (n - 1) / 26
	}
	return r
}
