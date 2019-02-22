/* https://leetcode.com/problems/reorder-log-files/description/
You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.
The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.
The digit-logs should be put in their original order.

Return the final order of the logs.

Example 1:

	Input: ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
	Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]

Note:

	0 <= logs.length <= 100
	3 <= logs[i].length <= 100
	logs[i] is guaranteed to have an identifier, and a word after the identifier.
*/

package lstring

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	res := make([]string, len(logs))
	tmp := make([]string, 0, len(logs))
	for i, j := len(logs)-1, len(logs)-1; i >= 0; i-- {
		flag := true
		for idx, char := range logs[i] {
			if char == ' ' && '0' <= logs[i][idx+1] && logs[i][idx+1] <= '9' {
				res[j] = logs[i]
				j--
				flag = false
				break
			}
		}
		if flag {
			tmp = append(tmp, logs[i])
		}

	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][strings.Index(tmp[i], " "):] < tmp[j][strings.Index(tmp[j], " "):]
	})
	copy(res[0:len(tmp)], tmp)
	return res
}
