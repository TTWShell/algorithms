package stack

import "sync"

type Stack struct {
	stack []interface{}
	len   int
	lock  sync.Mutex
}

func Constructor() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len == 0
}

func (s *Stack) Pop() (element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	element, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Push(element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := []interface{}{element}
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Top() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stack[0]
}
