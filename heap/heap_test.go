package heap

import "testing"

type Int int

func (i Int) LessThan(e Element) bool {
	return i < e.(Int)
}

func Test_Init(t *testing.T) {
	h := New(true)
	if h.len != 0 || h.Len() != 0 || len(h.heap) != 0 || h.isMin != true || !h.IsEmpty() {
		t.Fatal(h, "init MinHeap failed.")
	}

	h = New(false)
	if h.len != 0 || h.Len() != 0 || len(h.heap) != 0 || h.isMin != false || !h.IsEmpty() {
		t.Fatal(h, "init MaxHeap failed.")
	}
}

func Test_MinHeap(t *testing.T) {
	h := New(true)
	for _, num := range []Int{100, 90, 80, 5, 3, 1, 10, 20, 30, 8, 70} {
		h.Insert(num)
	}

	res := make([]Int, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract().(Int))
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			t.Fatal(res)
		}
	}
}

func Test_MaxHeap(t *testing.T) {
	h := New(false)
	for _, num := range []Int{100, 90, 80, 5, 3, 1, 10, 20, 30, 8, 70} {
		h.Insert(num)
	}

	res := make([]Int, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract().(Int))
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] < res[i+1] {
			t.Fatal(res)
		}
	}
}

func Test_ExtractPanic(t *testing.T) {
	h := New(true)
	defer func() {
		if r := recover(); r != "Empty heap, cannot Extract." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	h.Extract()
}

func Test_Peek(t *testing.T) {
	h := New(true)
	h.Insert(Int(100))

	if h.Peek().(Int) != Int(100) {
		t.Fatal(h.Peek().(Int))
	}

	h.Extract()

	defer func() {
		if r := recover(); r != "Empty heap, cannot Peek." {
			t.Fatal("Expected panic but err is:", r)
		}
	}()
	h.Peek()
}

func Test_precolate(t *testing.T) {
	h := New(true)
	for i, num := range []Int{100, 90, 80, 5, 3, 1, 10, 20, 30, 8, 70} {
		h.Insert(num)
		h.precolateUp(i)
		h.precolateDown(h.Len() - 1 - i)
	}

	res := make([]Int, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract().(Int))
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			t.Fatal(res)
		}
	}
}
