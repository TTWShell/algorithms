/* https://leetcode.com/problems/basic-calculator-ii/description/
Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.

You may assume that the given expression is always valid.

Some examples:
    "3+2*2" = 7
    " 3/2 " = 1
    " 3+5 / 2 " = 5
Note: Do not use the eval built-in library function.
*/

package leetcode

func calculate2(s string) int {
	numStack, operatorStack := []int{}, []byte{}
	for i := 0; i <= len(s); i++ {
		if i != len(s) {
			if tmp := s[i]; tmp == ' ' {
				continue
			} else if tmp == '+' || tmp == '-' || tmp == '*' || tmp == '/' {
				operatorStack = append(operatorStack, tmp)
				continue
			} else {
				num := 0
				for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
					num = num*10 + int(s[i]-'0')
				}
				i--
				numStack = append(numStack, num)
			}
		}

		if len(operatorStack) > 0 {
			switch operatorStack[len(operatorStack)-1] {
			case '*':
				numStack[len(numStack)-2] = numStack[len(numStack)-2] * numStack[len(numStack)-1]
			case '/':
				numStack[len(numStack)-2] = numStack[len(numStack)-2] / numStack[len(numStack)-1]
			default:
				continue
			}
			numStack = numStack[:len(numStack)-1]
			operatorStack = operatorStack[:len(operatorStack)-1]
		}
	}

	for len(operatorStack) > 0 {
		// only + or -
		switch operatorStack[0] {
		case '+':
			numStack[1] = numStack[0] + numStack[1]
		case '-':
			numStack[1] = numStack[0] - numStack[1]
		}
		numStack = numStack[1:]
		operatorStack = operatorStack[1:]
	}
	return numStack[0]
}
