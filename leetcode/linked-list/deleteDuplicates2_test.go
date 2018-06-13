package lll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	assert := assert.New(t)

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	assert.Equal(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}, deleteDuplicates2(head))

	head = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}
	assert.Equal(&ListNode{Val: 2, Next: &ListNode{Val: 3}}, deleteDuplicates2(head))
}
