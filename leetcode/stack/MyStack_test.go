package lstack

import "testing"

func Test_MyStack(t *testing.T) {
	stack := MyStackConstructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if r := stack.Empty(); r != false {
		t.Fatal("Empty:", r)
	}

	stack.Pop()
	if r := stack.Top(); r != 0 {
		t.Fatal("Top", r)
	}

	stack.Pop()
	stack.Pop()
	if r := stack.Empty(); r != true {
		t.Fatal("Empty:", r)
	}
}
