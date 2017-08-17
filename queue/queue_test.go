package queue

import "testing"

func Test_Queue(t *testing.T) {
	queue := Constructor()

	if !queue.IsEmpty() || queue.len != 0 || queue.Len() != 0 {
		t.Fail()
	}

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	if queue.queue[0] != 1 || queue.queue[1] != 2 || queue.queue[2] != 3 {
		t.Fatal(queue)
	}

	if queue.Len() != 3 {
		t.Fatal(queue)
	}

	if queue.DeQueue() != 1 || queue.Len() != 2 {
		t.Fatal(queue)
	}

	if queue.Peek() != 2 {
		t.Fatal(queue)
	}

	queue.DeQueue()
	queue.DeQueue()

	if queue.Len() != 0 {
		t.Fatal(queue)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		} else {
			if r != "Queue is empty, cannot DeQueue." {
				t.Error(r)
			}
		}
	}()
	queue.DeQueue()
}

func Test_QueuePeek(t *testing.T) {
	queue := Constructor()
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		} else {
			if r != "Queue is empty, cannot DeQueue." {
				t.Error(r)
			}
		}
	}()
	queue.Peek()
}
