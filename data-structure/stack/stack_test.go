package stack

import "testing"

func Test_Stack(t *testing.T) {
	stack := Constructor()

	if !stack.IsEmpty() || stack.len != 0 || stack.Len() != 0 {
		t.Error()
	}

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if stack.stack[0] != -2 || stack.stack[1] != 0 || stack.stack[2] != -3 {
		t.Fatal(stack.stack)
	}

	if stack.Len() != 3 {
		t.Fatal(stack)
	}

	if r := stack.IsEmpty(); r != false {
		t.Fatal("Empty:", r)
	}

	stack.Pop()
	if r := stack.Top(); r != 0 || stack.Len() != 2 {
		t.Fatal("Top", r, "stack", stack)
	}

	stack.Pop()
	stack.Pop()
	if r := stack.IsEmpty(); r != true {
		t.Fatal("Empty:", r)
	}
}
