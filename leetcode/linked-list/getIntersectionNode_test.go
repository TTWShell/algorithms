package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	assert := assert.New(t)

	intersect := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersect}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: intersect}}}
	assert.Equal(intersect, getIntersectionNode(headA, headB))
	assert.Equal(intersect, getIntersectionNode(headB, headA))

	headA = &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	headB = &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: intersect}}
	assert.Nil(getIntersectionNode(headA, headB))

	intersect = &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	headA = &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: intersect}}}
	headB = &ListNode{Val: 3, Next: intersect}
	assert.Equal(intersect, getIntersectionNode(headA, headB))
	assert.Equal(intersect, getIntersectionNode(headB, headA))
}
