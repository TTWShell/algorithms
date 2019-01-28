package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitListToParts(t *testing.T) {
	assert := assert.New(t)

	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	excepted := []*ListNode{&ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 3}, nil, nil}
	assert.Equal(excepted, splitListToParts(root, 5))
}
