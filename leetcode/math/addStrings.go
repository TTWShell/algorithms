/* https://leetcode.com/problems/add-strings/#/description
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

package leetcode

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m < n {
		temp := num1
		num1, num2 = num2, temp
		t := m
		m, n = n, t
	}

	for i := 0; i < m-n; i++ {
		num2 = "0" + num2
	}

	temp := make([]rune, m)
	carry := 0
	for i := m - 1; i >= 0; i-- {
		sum := carry + int(num1[i]-'0') + int(num2[i]-'0')
		if sum > 9 {
			carry = 1
			temp[i] = rune(sum - 10 + '0')
		} else {
			carry = 0
			temp[i] = rune(sum + '0')
		}
	}

	if carry == 1 {
		return "1" + string(temp)
	}
	return string(temp)
}
