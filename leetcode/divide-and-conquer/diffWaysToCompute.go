/* https://leetcode.com/problems/different-ways-to-add-parentheses/description/
Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.


Example 1
Input: "2-1-1".

	((2-1)-1) = 0
	(2-(1-1)) = 2

Output: [0, 2]

Example 2
Input: "2*3-4*5"

	(2*(3-(4*5))) = -34
	((2*3)-(4*5)) = -14
	((2*(3-4))*5) = -10
	(2*((3-4)*5)) = -10
	(((2*3)-4)*5) = 10

Output: [-34, -14, -10, -10, 10]
*/

package leetcode

import (
	"bytes"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	tmp := []string{}
	for i := 0; i < len(input); i++ {
		if letter := input[i]; letter == '+' || letter == '-' || letter == '*' {
			tmp = append(tmp, string(letter))
		} else {
			var num bytes.Buffer
			for ; i < len(input) && '0' <= input[i] && input[i] <= '9'; i++ {
				num.WriteByte(input[i])
			}
			i--
			tmp = append(tmp, num.String())
		}
	}

	var helper func(input []string) []int
	helper = func(input []string) []int {
		if len(input) == 1 {
			num, _ := strconv.Atoi(input[0])
			return []int{num}
		} else if len(input) == 3 {
			num1, _ := strconv.Atoi(input[0])
			num2, _ := strconv.Atoi(input[2])
			switch input[1] {
			case "+":
				return []int{num1 + num2}
			case "-":
				return []int{num1 - num2}
			case "*":
				return []int{num1 * num2}
			}
		}

		res := []int{}
		for i := 0; i < len(input)-1; i += 2 {
			for _, num1 := range helper(input[:i+1]) {
				for _, num2 := range helper(input[i+2:]) {
					switch input[i+1] {
					case "*":
						res = append(res, num1*num2)
					case "+":
						res = append(res, num1+num2)
					case "-":
						res = append(res, num1-num2)
					}
				}
			}
		}
		return res
	}

	return helper(tmp)
}
