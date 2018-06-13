package lll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderList(t *testing.T) {
	assert := assert.New(t)

	reorderList(nil) // test empty input can work

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	excepted := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	reorderList(head)
	assert.Equal(excepted, head)
}
