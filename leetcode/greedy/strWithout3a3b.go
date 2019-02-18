/* https://leetcode.com/problems/string-without-aaa-or-bbb/description/
Given two integers A and B, return any string S such that:

S has length A + B and contains exactly A 'a' letters, and exactly B 'b' letters;
The substring 'aaa' does not occur in S;
The substring 'bbb' does not occur in S.


Example 1:

	Input: A = 1, B = 2
	Output: "abb"
	Explanation: "abb", "bab" and "bba" are all correct answers.

Example 2:

	Input: A = 4, B = 1
	Output: "aabaa"

Note:

	0 <= A <= 100
	0 <= B <= 100
	It is guaranteed such an S exists for the given A and B.
*/

package lgreedy

import "bytes"

func strWithout3a3b(A int, B int) string {
	less, more := byte('a'), byte('b') // A < B
	if A > B {
		A, B = B, A
		less, more = more, less
	}

	tmp := string([]byte{more, more, less})
	var res bytes.Buffer
	for B > A && A > 0 {
		res.WriteString(tmp)
		A--
		B -= 2
	}
	for A > 0 || B > 0 {
		res.WriteByte(more)
		B--
		if A > 0 {
			res.WriteByte(less)
			A--
		}
	}
	return res.String()
}
