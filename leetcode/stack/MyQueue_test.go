package lstack

import "testing"

func Test_MyQueue(t *testing.T) {
	myqueue := MyQueueConstructor()
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)

	if r := myqueue.Empty(); r != false {
		t.Fatal("Empty:", r)
	}

	if r := myqueue.Peek(); r != 1 {
		t.Fatal("Peek:", r)
	}

	myqueue.Pop()
	myqueue.Pop()
	myqueue.Pop()
	if r := myqueue.Empty(); r != true {
		t.Fatal("Empty:", r)
	}
}
