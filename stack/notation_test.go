package stack

import (
	"strings"
	"testing"
)

func Test_isNum(t *testing.T) {
	if r := isNum("2"); r != true {
		t.Fatal(r)
	}

	if r := isNum("0.023231"); r != true {
		t.Fatal(r)
	}
}

func Test_IN2RPN(t *testing.T) {
	infix := "4.99 * 1.06 + 5.99 + 6.99 * 1.06"
	reversePolish := "4.99 1.06 * 5.99 + 6.99 1.06 * +"

	if r := IN2RPN(strings.Split(infix, " ")); strings.Join(r, " ") != reversePolish {
		t.Fatal(r)
	}

	infix = "( ( 4.99 * 1.06 ) + 5.99 ) + ( 6.99 * 1.06 )"
	if r := IN2RPN(strings.Split(infix, " ")); strings.Join(r, " ") != reversePolish {
		t.Fatal(r)
	}
}
