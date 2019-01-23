package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	assert := assert.New(t)

	toArray := func(head *ListNode) []int {
		res := []int{}
		for ; head != nil; head = head.Next {
			res = append(res, head.Val)
		}
		return res
	}
	head := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}}
	node := head.Next
	deleteNode(node)

	assert.Equal([]int{4, 1, 9}, toArray(head))
}
