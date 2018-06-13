package lstack

import "testing"

func Test_MinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if r := stack.GetMin(); r != -3 {
		t.Fatal("GetMin:", r)
	}

	stack.Pop()
	if r := stack.Top(); r != 0 {
		t.Fatal("Top", r)
	}

	if r := stack.GetMin(); r != -2 {
		t.Fatal("GetMin:", r)
	}
}
