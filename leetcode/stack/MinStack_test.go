package leetcode

import "testing"

func Test_MinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if r := stack.GetMin(); r != -3 {
		t.Log("GetMin:", r)
		t.Fail()
	}

	stack.Pop()
	if r := stack.Top(); r != 0 {
		t.Log("Top", r)
		t.Fail()
	}

	if r := stack.GetMin(); r != -2 {
		t.Log("GetMin:", r)
		t.Fail()
	}
}
