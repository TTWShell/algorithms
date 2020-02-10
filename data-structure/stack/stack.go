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

	element = s.stack[s.len-1]
	s.stack = s.stack[:s.len-1]
	s.len--
	return
}

func (s *Stack) Push(element interface{}) {
	s.Lock()
	defer s.Unlock()

	s.stack = append(s.stack, element)
	s.len++
}

func (s *Stack) Top() interface{} {
	s.RLock()
	defer s.RUnlock()

	return s.stack[s.len-1]
}
