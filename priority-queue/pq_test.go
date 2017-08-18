package pq

import (
	"math/rand"
	"testing"
)

var elList = []*Element{
	NewElement(100, 5), NewElement(90, 6), NewElement(70, 7),
	NewElement(4, 8), NewElement(5, 9), NewElement(6, 10),
	NewElement(10, 11), NewElement(15, 12), NewElement(80, 13),
}

func TestMinPriorityQueue(t *testing.T) {
	pq := MinPQConstructor()
	for _, el := range elList {
		pq.Insert(*el)
	}

	res := make([]Element, 0)
	for pq.Len() > 0 {
		res = append(res, pq.Extract())
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i].Priority > res[i+1].Priority {
			t.Fatal(res)
		}
	}
}
func Test_MaxPriorityQueue(t *testing.T) {
	pq := MaxPQConstructor()
	for _, el := range elList {
		pq.Insert(*el)
	}

	res := make([]Element, 0)
	for pq.Len() > 0 {
		res = append(res, pq.Extract())
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
	pq := MaxPQConstructor()
	el := NewElement(100, 5)
	pq.Insert(*el)

	if pq.Peek() != *el {
		t.Fatal(pq.Peek())
	}

	pq.Extract()

	defer func() {
		if r := recover(); r != "Empty pq, cannot Peek." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	pq.Peek()
}

func Test_ChangePriority(t *testing.T) {
	pq := MaxPQConstructor()
	for _, el := range elList {
		pq.Insert(*el)
	}

	old := pq.Peek()
	if old.Value != 80 {
		t.Fatal(old)
	}

	pq.ChangePriority(0, -1)
	pq.ChangePriority(1, 100)

	if pq.Extract().Priority != 100 {
		t.Fatal(pq)
	}

	for pq.Len() != 0 {
		if pq.Len() == 1 && (pq.Peek().Value != old.Value || pq.Peek().Priority != -1) {
			t.Fatal(pq)
		}
		pq.Extract()
	}
}

func Test_ChangePriority2(t *testing.T) {
	pq := MaxPQConstructor()
	for _, el := range elList {
		pq.Insert(*el)
		pq.ChangePriority(0, rand.Int())
	}

	res := make([]Element, 0)
	for pq.Len() > 0 {
		res = append(res, pq.Extract())
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i].Priority < res[i+1].Priority {
			t.Fatal(res)
		}
	}
}
