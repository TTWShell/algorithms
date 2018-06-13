package lll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	assert := assert.New(t)

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}}
	assert.Equal(expected, reverseBetween(head, 2, 4))
	assert.Equal(head, reverseBetween(head, 3, 3))
}
