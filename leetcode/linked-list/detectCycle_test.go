package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(detectCycle(nil))

	loop := &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}
	head := &ListNode{Val: 3, Next: loop}
	assert.Nil(detectCycle(head))

	next := loop
	for next.Next != nil {
		next = next.Next
	}
	next.Next = loop
	head = &ListNode{Val: 3, Next: loop}
	assert.Equal(loop, detectCycle(head))
}
