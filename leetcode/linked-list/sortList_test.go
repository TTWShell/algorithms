package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortList(t *testing.T) {
	assert := assert.New(t)

	var head *ListNode
	assert.Equal(head, sortList(head))
	assert.Equal(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		sortList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}))
	assert.Equal(&ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		sortList(&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}))

	assert.Equal(&ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: &ListNode{Val: 10}}}},
		sortList(&ListNode{Val: 10, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8}}}}))
}
