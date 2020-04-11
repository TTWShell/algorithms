/* https://leetcode.com/problems/expression-add-operators/
Given a string that contains only digits 0-9 and a target value,
return all possibilities to add binary operators (not unary)
+, -, or * between the digits so they evaluate to the target value.

Example 1:

	Input: num = "123", target = 6
	Output: ["1+2+3", "1*2*3"]

Example 2:

	Input: num = "232", target = 8
	Output: ["2*3+2", "2+3*2"]

Example 3:

	Input: num = "105", target = 5
	Output: ["1*0+5","10-5"]

Example 4:

	Input: num = "00", target = 0
	Output: ["0+0", "0-0", "0*0"]

Example 5:

	Input: num = "3456237490", target = 9191
	Output: []
*/

package ldc

import "strings"

func addOperators(num string, target int) []string {
	res := make([]string, 0)
	if len(num) == 0 {
		return res
	}

	var dfs func(res []string, path []rune, length int, left int, cur int, digits []rune, pos int, target int) []string
	dfs = func(res []string, path []rune, length int, left int, cur int, digits []rune, pos int, target int) []string {
		if pos == len(digits) {
			if left+cur == target {
				addpath := string(path)
				addpath = strings.Replace(addpath, " ", "", -1)
				res = append(res, string(addpath))
			}
			return res
		}

		n := 0
		j := length + 1
		for i := pos; i < len(digits); i++ {
			n = n*10 + int(digits[i]) - '0'
			path[j] = digits[i]
			j++
			if j < len(path) {
				path[j] = ' '
			}

			path[length] = '+'
			res = dfs(res, path, j, left+cur, n, digits, i+1, target)

			path[length] = '-'
			res = dfs(res, path, j, left+cur, -n, digits, i+1, target)

			path[length] = '*'
			res = dfs(res, path, j, left, cur*n, digits, i+1, target)

			if digits[pos] == '0' {
				break
			}
		}
		return res
	}

	path := make([]rune, len(num)*2-1)
	digits := []rune(num)
	n := 0
	for i := 0; i < len(digits); i++ {
		n = n*10 + int(digits[i]-'0')
		path[i] = digits[i]
		if i < len(path) {
			path[i+1] = ' '
		}
		res = dfs(res, path, i+1, 0, n, digits, i+1, target)
		if n == 0 {
			break
		}
	}
	return res
}
