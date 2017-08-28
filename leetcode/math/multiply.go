/* https://leetcode.com/problems/multiply-strings/description/
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.

Note:

    1. The length of both num1 and num2 is < 110.
    2. Both num1 and num2 contains only digits 0-9.
    3. Both num1 and num2 does not contain any leading zero.
    4. You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

package leetcode

func multiply(num1 string, num2 string) string {
	sum := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		var carry, bit byte
		for j := len(num2) - 1; j >= 0; j-- {
			bit = byte(num1[i]-'0')*byte(num2[j]-'0') + carry
			if sum[i+j+1] != 0 {
				bit += byte(sum[i+j+1] - '0')
			}
			sum[i+j+1] = bit%10 + '0'
			carry = bit / 10
		}
		if carry != 0 {
			sum[i] = carry + '0'
		}
	}

	for i := 0; i < len(sum); i++ {
		if sum[i] != 0 && sum[i]-'0' > 0 {
			return string(sum[i:])
		}
	}
	return "0"
}
