package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers2(t *testing.T) {
	assert := assert.New(t)

	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	excepted := &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 7}}}}
	assert.Equal(excepted, addTwoNumbers2(l1, l2))

	var empty *ListNode
	assert.Equal(l1, addTwoNumbers2(l1, empty))
	assert.Equal(l1, addTwoNumbers2(empty, l1))

	l1 = &ListNode{Val: 1}
	l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	excepted = &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}
	assert.Equal(excepted, addTwoNumbers2(l1, l2))
}
