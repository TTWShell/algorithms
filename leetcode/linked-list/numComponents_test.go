package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numComponents(t *testing.T) {
	assert := assert.New(t)

	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	G := []int{0, 1, 3}
	assert.Equal(2, numComponents(head, G))

	head = &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}
	G = []int{0, 3, 1, 4}
	assert.Equal(2, numComponents(head, G))
}
