/* https://leetcode.com/problems/repeated-string-match/description/

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/

package leetcode

import (
	"bytes"
	"math"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	times := int(math.Ceil(float64(len(B)) / float64(len(A))))

	var buffer bytes.Buffer
	for i := 0; i < times-1; i++ {
		buffer.WriteString(A)
	}

	for i := 0; i < 2; i++ {
		buffer.WriteString(A)
		if strings.Contains(buffer.String(), B) {
			return times + i
		}
	}
	return -1
}
