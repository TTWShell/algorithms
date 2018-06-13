/* https://leetcode.com/problems/number-of-segments-in-a-string/#/description
Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any non-printable characters.

Example:

    Input: "Hello, my name is John"
    Output: 5
*/

package lstring

func countSegments(s string) int {
	count, char := 0, 0
	for _, r := range s {
		if r != rune(' ') {
			if char == 0 {
				count++
			}
			char++
		} else {
			char = 0
		}
	}
	return count
}
