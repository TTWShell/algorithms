/* https://leetcode.com/problems/find-median-from-data-stream/
Median is the middle value in an ordered integer list. If the size of the list is even,
there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.


Example:

	addNum(1)
	addNum(2)
	findMedian() -> 1.5
	addNum(3)
	findMedian() -> 2


Follow up:

	If all integer numbers from the stream are between 0 and 100, how would you optimize it?
	If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?
*/

package ldesign

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	size := len(*h)
	v := (*h)[size-1]
	*h = (*h)[0 : size-1]
	return v
}

func (h *IntHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *IntHeap) Peak() interface{} {
	return (*h)[0]
}

type MedianFinder struct {
	minHeap *IntHeap
	maxHeap *IntHeap
	size    int
}

/** initialize your data structure here. */
func MedianFinderConstructor() MedianFinder {
	minHeap, maxHeap := &IntHeap{}, &IntHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{minHeap: minHeap, maxHeap: maxHeap}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 {
		heap.Push(this.minHeap, num)
		return
	}
	if num >= this.minHeap.Peak().(int) {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, -num)
	}

	if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
	if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(this.minHeap.Peak().(int)-this.maxHeap.Peak().(int)) / 2
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.Peak().(int))
	}
	return float64(-this.maxHeap.Peak().(int))
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
