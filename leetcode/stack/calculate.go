/* https://leetcode.com/problems/basic-calculator/
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ),
the plus + or minus sign -, non-negative integers and empty spaces .

Example 1:

	Input: "1 + 1"
	Output: 2

Example 2:

	Input: " 2-1 + 2 "
	Output: 3

Example 3:

	Input: "(1+(4+5+2)-3)+(6+8)"
	Output: 23

Note:
	You may assume that the given expression is always valid.
	Do not use the eval built-in library function.
*/

package lstack

import (
	"strconv"
)

func calculate(s string) int {
	stackNum := []int{}
	stackSymbol := []byte{}

	helper := func() {
		if len(stackSymbol) == 0 {
			return
		}

		symbol := stackSymbol[len(stackSymbol)-1]
		stackSymbol = stackSymbol[:len(stackSymbol)-1]
		if symbol == '(' {
			return
		}

		num1, num2 := stackNum[len(stackNum)-2], stackNum[len(stackNum)-1]
		var tmp int
		if symbol == '+' {
			tmp = num1 + num2
		} else {
			tmp = num1 - num2
		}
		stackNum = stackNum[:len(stackNum)-2]
		stackNum = append(stackNum, tmp)
	}

	numArray := []byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if s[i] == ' ' {
			continue
		}
		if '0' <= char && char <= '9' {
			numArray = append(numArray, char)
			continue
		}

		if len(numArray) != 0 {
			num, _ := strconv.Atoi(string(numArray))
			numArray = []byte{}
			stackNum = append(stackNum, num)
			helper()
		}

		if char != ')' {
			stackSymbol = append(stackSymbol, char)
		} else {
			helper()
		}
	}
	if len(numArray) != 0 {
		num, _ := strconv.Atoi(string(numArray))
		stackNum = append(stackNum, num)
	}
	helper()

	return stackNum[0]
}
