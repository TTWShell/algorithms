package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	assert := assert.New(t)

	assert.False(hasCycle(nil))

	loop := &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}
	head := &ListNode{Val: 3, Next: loop}
	assert.False(hasCycle(head))

	next := loop
	for next.Next != nil {
		next = next.Next
	}
	next.Next = loop
	head = &ListNode{Val: 3, Next: loop}
	assert.True(hasCycle(head))
}
