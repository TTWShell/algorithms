/* https://leetcode.com/problems/additive-number/description/
Additive number is a string whose digits can form additive sequence.

A valid additive sequence should contain at least three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

For example:

    "112358" is an additive number because the digits can form an additive sequence: 1, 1, 2, 3, 5, 8.

    1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
    "199100199" is also an additive number, the additive sequence is: 1, 99, 100, 199.
    1 + 99 = 100, 99 + 100 = 199

Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.

Given a string containing only digits '0'-'9', write a function to determine if it's an additive number.

Follow up:
How would you handle overflow for very large input integers?
*/

package leetcode

import "bytes"

func isAdditiveNumber(num string) bool {
	addStrings := func(num1 string, num2 string) string {
		// https://leetcode.com/problems/add-strings/description/
		m, n := len(num1), len(num2)
		if m < n {
			temp := num1
			num1, num2 = num2, temp
			t := m
			m, n = n, t
		}

		var leadingZero bytes.Buffer
		for i := 0; i < m-n; i++ {
			leadingZero.WriteString("0")
		}
		leadingZero.WriteString(num2)
		num2 = leadingZero.String()

		temp := make([]rune, m, m)
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

	for i := 1; i <= len(num)/2; i++ {
		for j := i + 1; j < len(num); j++ {
			start, mid, end := 0, i, j
			for {
				if num[mid] == '0' && end-mid != 1 {
					break
				}
				sum := addStrings(num[start:mid], num[mid:end])
				tmpEnd := end + len(sum)
				if tmpEnd > len(num) || sum != num[end:tmpEnd] {
					break
				}
				start, mid, end = mid, end, tmpEnd
				if end == len(num) {
					return true
				}
			}
		}
	}
	return false
}
