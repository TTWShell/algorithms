package stack

import "sync"

type Stack struct {
	stack []interface{}
	len   int
	sync.RWMutex
}

func Constructor() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	s.RLock()
	defer s.RUnlock()

	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.len == 0
}

func (s *Stack) Pop() (element interface{}) {
	s.Lock()
	defer s.Unlock()

	element, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Push(element interface{}) {
	s.Lock()
	defer s.Unlock()

	prepend := []interface{}{element}
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Top() interface{} {
	s.RLock()
	defer s.RUnlock()

	return s.stack[0]
}
