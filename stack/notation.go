package stack

import (
	"fmt"
	"strconv"
)

var priority = map[string]int{
	"(": 1, ")": 1,
	"*": 3, "/": 3,
	"+": 4, "-": 4,
}

func isNum(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return false
	}
	return true
}

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func isBracket(s string) bool {
	if s == "(" || s == ")" {
		return true
	}
	return false
}

func IN2PN(infix []string) (polish []string) {
	s1, s2 := Constructor(), Constructor()

	for i := len(infix) - 1; i >= 0; i-- {
		s := infix[i]
		if isNum(s) {
			s2.Push(s)
		} else if isOperator(s) {
			if s1.IsEmpty() || s1.Top().(string) == ")" ||
				priority[s] <= priority[s1.Top().(string)] {
				s1.Push(s)
			} else {
				s2.Push(s1.Pop())
				i++
			}
		} else if isBracket(s) {
			if s == ")" {
				s1.Push(s)
			} else {
				for !s1.IsEmpty() && s1.Top().(string) != ")" {
					s2.Push(s1.Pop())
				}
				if s1.Top().(string) == ")" {
					s1.Pop()
				}
			}
		} else {
			panic(fmt.Sprintf("error input: %s", s))
		}
	}

	for !s1.IsEmpty() {
		s2.Push(s1.Pop())
	}

	res, index := make([]string, s2.Len(), s2.Len()), 0
	for !s2.IsEmpty() {
		res[index] = s2.Pop().(string)
		index++
	}
	return res
}

func IN2RPN(infix []string) (reversePolish []string) {
	s1, s2 := Constructor(), Constructor()

	for i := 0; i < len(infix); i++ {
		s := infix[i]
		if isNum(s) {
			s2.Push(s)
		} else if isOperator(s) {
			if s1.IsEmpty() || s1.Top().(string) == "(" ||
				priority[s] < priority[s1.Top().(string)] {
				s1.Push(s)
			} else {
				s2.Push(s1.Pop())
				i--
			}
		} else if isBracket(s) {
			if s == "(" {
				s1.Push(s)
			} else {
				for !s1.IsEmpty() && s1.Top().(string) != "(" {
					s2.Push(s1.Pop())
				}
				if s1.Top().(string) == "(" {
					s1.Pop()
				}
			}
		} else {
			panic(fmt.Sprintf("error input: %s", s))
		}
	}

	for !s1.IsEmpty() {
		s2.Push(s1.Pop())
	}

	res, index := make([]string, s2.Len(), s2.Len()), s2.Len()-1
	for !s2.IsEmpty() {
		res[index] = s2.Pop().(string)
		index--
	}
	return res
}
