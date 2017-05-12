/* https://leetcode.com/problems/add-binary/#/description
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

package leetcode

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m < n {
		m, n = n, m
		a, b = b, a
		for i := m - n; i > 0; i-- {
			b = "0" + b
		}
	}

	plus, r := 0, ""
	for i := m - 1; i != -1; i-- {
		ans := int(a[i]) + int(b[i]) + plus
		if ans >= 98 {
			r = string(ans-'2') + r
			plus = 1
		} else {
			r = string(ans-'0') + r
			plus = 0
		}
	}
	if plus == 1 {
		r = "1" + r
	}
	return r
}
