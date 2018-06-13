package lll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partition(t *testing.T) {
	assert := assert.New(t)

	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	excepted := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}}
	assert.Equal(excepted, partition(head, 3))

	head = &ListNode{Val: 1}
	excepted = &ListNode{Val: 1}
	assert.Equal(excepted, partition(head, 2))
	assert.Equal(excepted, partition(head, 0))
}
