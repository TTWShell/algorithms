/* https://leetcode.com/problems/decode-ways/description/
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given an encoded message containing digits, determine the total number of ways to decode it.

For example,
Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

The number of ways decoding "12" is 2.
*/

package leetcode

func numDecodings(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	dp := make([]int, length+1, length+1)
	dp[length] = 1
	if letter := s[length-1]; '0' < letter && letter <= '9' {
		dp[length-1] = 1
	}

	for i := length - 2; i >= 0; i-- {
		letter := s[i]
		if letter < '1' || letter > '9' {
			continue
		}
		if (letter-'0')*10+s[i+1]-'0' <= 26 {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
