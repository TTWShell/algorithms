package heap

import "testing"

type Int int

func (i Int) LessThan(e Element) bool {
	return i < e.(Int)
}

func Test_MinHeap(t *testing.T) {
	h := MinHeapConstructor()

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
	h := MaxHeapConstructor()

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