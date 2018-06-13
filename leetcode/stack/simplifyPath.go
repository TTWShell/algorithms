/*https://leetcode.com/problems/simplify-path/description/
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

Corner Cases:
    Did you consider the case where path = "/../"?
    In this case, you should return "/".
    Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
    In this case, you should ignore redundant slashes and return "/home/foo".
*/

package lstack

import (
	"github.com/TTWShell/algorithms/data-structure/stack" // need copy stack.go when run in leetcode online
	"strings"
)

func simplifyPath(path string) string {
	s := stack.Constructor()
	paths := strings.Split(path, "/")

	for _, sub := range paths {
		if sub == "." || sub == "" {
			continue
		} else if sub == ".." {
			if !s.IsEmpty() {
				s.Pop()
			}
		} else {
			s.Push(sub)
		}
	}

	if length := s.Len(); length > 0 {
		res, idx := make([]string, length, length), length-1
		for !s.IsEmpty() {
			res[idx] = s.Pop().(string)
			idx--
		}
		return "/" + strings.Join(res, "/")
	}
	return "/"
}

/*
type Stack struct {
	stack []interface{}
	len   int
}

func Constructor() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Pop() (element interface{}) {
	element, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Push(element interface{}) {
	prepend := []interface{}{element}
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Top() interface{} {
	return s.stack[0]
}
*/
