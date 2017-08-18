package pq

import (
	"github.com/TTWShell/algorithms/heap"
)

type Element struct {
	Value    interface{}
	Priority int
}

func NewElement(value interface{}, priority int) (e *Element) {
	return &Element{
		Value:    value,
		Priority: priority,
	}
}

func (e Element) LessThan(hel heap.Element) bool {
	return e.Priority < hel.(Element).Priority
}

type PQ struct {
	data heap.Heap
}

func MaxPQConstructor() (pq *PQ) {
	return &PQ{data: *heap.MaxHeapConstructor()}
}

func MinPQConstructor() (pq *PQ) {
	return &PQ{data: *heap.MinHeapConstructor()}
}

func (pq *PQ) Len() int {
	return pq.data.Len()
}

func (pq *PQ) Insert(el Element) {
	pq.data.Insert(heap.Element(el))
}

func (pq *PQ) Extract() (el Element) {
	return pq.data.Extract().(Element)
}
