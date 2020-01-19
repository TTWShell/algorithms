package pq

import (
	"sync"
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

func (e Element) GreaterThan(hel Element) bool {
	return e.Priority > hel.Priority
}

type PQ struct {
	sync.RWMutex
	heap []Element
	len  int
}

func PQConstructor() (pq *PQ) {
	return &PQ{heap: make([]Element, 0)}
}

func (pq *PQ) Len() int {
	pq.RLock()
	defer pq.RUnlock()

	return pq.len
}

func (pq *PQ) Insert(el Element) {
	pq.Lock()
	defer pq.Unlock()

	pq.heap = append(pq.heap, el)
	pq.len++
	pq.precolateUp(pq.len - 1)
}

func (pq *PQ) Extract() (el Element) {
	pq.Lock()
	defer pq.Unlock()

	if pq.len == 0 {
		panic("Empty pq, cannot Extract.")
	}

	pq.len--
	el, pq.heap[0] = pq.heap[0], pq.heap[pq.len]
	pq.heap = pq.heap[:pq.len]
	pq.precolateDown(0)
	return
}

func (pq *PQ) Peek() (el Element) {
	pq.RLock()
	defer pq.RUnlock()

	if pq.len == 0 {
		panic("Empty pq, cannot Peek.")
	}
	return pq.heap[0]
}

// func (pq *PQ) ChangePriority(index, priority int) {
// 	pq.Lock()
// 	defer pq.Unlock()

// 	oldEl := pq.heap[index]
// 	pq.heap[index].Priority = priority
// 	if pq.less(Element{Priority: priority}, oldEl) {
// 		pq.precolateUp(index)
// 	} else {
// 		pq.precolateDown(index)
// 	}
// }

func (pq *PQ) precolateUp(index int) {
	// 上滤，新元素在堆中上滤直到找出正确位置
	needUp, parent := index, (index-1)>>1
	for needUp > 0 && pq.heap[needUp].GreaterThan(pq.heap[parent]) {
		pq.heap[parent], pq.heap[needUp] = pq.heap[needUp], pq.heap[parent]
		// needUp, parent = parent, (parent-1)>>1
		needUp = parent
		parent = (parent - 1) >> 1
		// fmt.Println(needUp, parent)
	}
}

func (pq *PQ) precolateDown(index int) {
	// 下滤
	needDown, child := index, (index<<1)+1
	for needDown < pq.len && child < pq.len {
		// find max(leftChild, rightChild)
		if child+1 < pq.len && pq.heap[child+1].GreaterThan(pq.heap[child]) {
			child++
		}

		if pq.heap[needDown].GreaterThan(pq.heap[child]) {
			break
		}

		pq.heap[needDown], pq.heap[child] = pq.heap[child], pq.heap[needDown]
		needDown = child
		child = (needDown << 1) + 1
		// needDown, child = child, (needDown<<1)+1
		// fmt.Println(needDown, child)
	}
}
