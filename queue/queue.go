package queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	sync.RWMutex
}

func Constructor() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0

	return queue
}

func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()

	return q.len
}

func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return q.len == 0
}

func (q *Queue) DeQueue() (element interface{}) {
	// 删除队头元素并返回
	q.Lock()
	defer q.Unlock()

	if q.len <= 0 {
		panic("Queue is empty, cannot DeQueue.")
	}

	element, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) EnQueue(element interface{}) {
	// 插入元素到队尾
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, element)
	q.len++
}

func (q *Queue) Peek() interface{} {
	// 获取队头元素
	q.RLock()
	defer q.RUnlock()

	if q.len <= 0 {
		panic("Queue is empty, cannot DeQueue.")
	}
	return q.queue[0]
}
