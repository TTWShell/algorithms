package heapsort

import (
	"github.com/TTWShell/algorithms/heap"
)

type Int int

func (i Int) LessThan(e heap.Element) bool {
	return i < e.(Int)
}

func HeapSort(arr []int) {
	h := heap.MinHeapConstructor()
	for i := 0; i < len(arr); i++ {
		h.Insert(Int(arr[i]))
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = int(h.Extract().(Int))
	}
}
