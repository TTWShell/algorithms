package pq

import "sync"

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

func (e Element) LessThan(hel Element) bool {
	return e.Priority < hel.Priority
}

type PQ struct {
	sync.Mutex
	heap  []Element
	len   int
	isMin bool
}

func MinPQConstructor() (pq *PQ) {
	return &PQ{heap: make([]Element, 0), isMin: true}
}

func MaxPQConstructor() (pq *PQ) {
	return &PQ{heap: make([]Element, 0), isMin: false}
}

func (pq *PQ) Len() int {
	pq.Lock()
	defer pq.Unlock()

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

	el, pq.heap[0] = pq.heap[0], pq.heap[pq.len-1]
	pq.heap = pq.heap[:pq.len-1]
	pq.len--
	pq.precolateDown(0)
	return
}

func (pq *PQ) Peek() (el Element) {
	pq.Lock()
	defer pq.Unlock()

	if pq.len == 0 {
		panic("Empty pq, cannot Peek.")
	}
	return pq.heap[0]
}

func (pq *PQ) ChangePriority(index, priority int) {
	pq.Lock()
	defer pq.Unlock()

	oldEl := pq.heap[index]
	pq.heap[index].Priority = priority
	if pq.less(Element{Priority: priority}, oldEl) {
		pq.precolateUp(index)
	} else {
		pq.precolateDown(index)
	}
}

func (pq *PQ) precolateUp(index int) {
	// 上滤，新元素在堆中上滤直到找出正确位置
	needUp, parent := index, index>>1
	for needUp > 0 && pq.less(pq.heap[needUp], pq.heap[parent]) {
		pq.heap[parent], pq.heap[needUp] = pq.heap[needUp], pq.heap[parent]
		needUp, parent = parent, parent>>1
	}
}

func (pq *PQ) precolateDown(index int) {
	// 下滤
	needDown, child := index, index<<1+1
	for needDown < pq.len && child < pq.len {
		// find min(leftChild, rightChild)
		if child+1 < pq.len && pq.less(pq.heap[child+1], pq.heap[child]) {
			child++
		}

		if pq.less(pq.heap[needDown], pq.heap[child]) {
			break
		}

		pq.heap[needDown], pq.heap[child] = pq.heap[child], pq.heap[needDown]
		needDown, child = child, needDown<<1+1
	}
}

func (pq *PQ) less(a, b Element) bool {
	if pq.isMin {
		return a.LessThan(b)
	} else {
		return b.LessThan(a)
	}
}
