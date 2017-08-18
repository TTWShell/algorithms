package pq

import "testing"

var elList = []*Element{
	NewElement(100, 5), NewElement(90, 6), NewElement(70, 7),
	NewElement(4, 8), NewElement(5, 9), NewElement(6, 10),
	NewElement(10, 11), NewElement(15, 12), NewElement(80, 13),
}

func TestMinPriorityQueue(t *testing.T) {
	h := MinPQConstructor()
	for _, el := range elList {
		h.Insert(*el)
	}

	res := make([]Element, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract())
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i].Priority > res[i+1].Priority {
			t.Fatal(res)
		}
	}
}
func Test_MaxPriorityQueue(t *testing.T) {
	h := MaxPQConstructor()
	for _, el := range elList {
		h.Insert(*el)
	}

	res := make([]Element, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract())
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i].Priority < res[i+1].Priority {
			t.Fatal(res)
		}
	}
}

func Test_ExtractPanic(t *testing.T) {
	pq := MinPQConstructor()
	defer func() {
		if r := recover(); r != "Empty pq, cannot Extract." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	pq.Extract()
}

func Test_Peek(t *testing.T) {
	h := MaxPQConstructor()
	el := NewElement(100, 5)
	h.Insert(*el)

	if h.Peek() != *el {
		t.Fatal(h.Peek())
	}

	h.Extract()

	defer func() {
		if r := recover(); r != "Empty pq, cannot Peek." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	h.Peek()
}
