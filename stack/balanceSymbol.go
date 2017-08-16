package stack

func IsBalanceSymbol(s string) bool {
	stack := Constructor()
	maps := map[rune]rune{']': '[', ')': '(', '}': '{'}

	for _, r := range s {
		if r == '[' || r == '(' || r == '{' {
			stack.Push(r)
		} else if r == ']' || r == ')' || r == '}' {
			if stack.IsEmpty() {
				return false
			} else if top := stack.Pop(); top != maps[r] {
				return false
			}
		}
	}

	if !stack.IsEmpty() {
		return false
	}
	return true
}
