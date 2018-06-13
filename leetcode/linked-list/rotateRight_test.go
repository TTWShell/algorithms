package lll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(rotateRight(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1), &ListNode{Val: 2, Next: &ListNode{Val: 1}})

	assert.Equal(rotateRight(&ListNode{Val: 1}, 99), &ListNode{Val: 1})

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	output := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}
	assert.Equal(rotateRight(head, 2), output)
}
