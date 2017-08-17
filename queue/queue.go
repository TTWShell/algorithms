package queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func Constructor() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len == 0
}

func (q *Queue) DeQueue() (element interface{}) {
	// 删除队头元素并返回
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len <= 0 {
		panic("Queue is empty, cannot DeQueue.")
	}

	element, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) EnQueue(element interface{}) {
	// 插入元素到队尾
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, element)
	q.len++
}

func (q *Queue) Peek() interface{} {
	// 获取队头元素
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len <= 0 {
		panic("Queue is empty, cannot DeQueue.")
	}
	return q.queue[0]
}
