package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	assert := assert.New(t)

	middle := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: middle}}
	assert.Equal(middle, middleNode(head))

	middle = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: middle}}}
	assert.Equal(middle, middleNode(head))
}
