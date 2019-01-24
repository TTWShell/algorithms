package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oddEvenList(t *testing.T) {
	assert := assert.New(t)

	list2array := func(head *ListNode) []int {
		res := []int{}
		for cur := head; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}

	array2list := func(array []int) *ListNode {
		var head, cur *ListNode
		for i := range array {
			node := &ListNode{Val: array[i]}
			if head == nil {
				head = node
			} else {
				cur.Next = node
			}
			cur = node
		}
		return head
	}

	head := array2list([]int{1, 2, 3, 4, 5})
	assert.Equal([]int{1, 3, 5, 2, 4}, list2array(oddEvenList(head)))

	// head = array2list([]int{1, 2, 3, 4, 5, 6})
	// assert.Equal([]int{1, 3, 5, 2, 4, 6}, list2array(oddEvenList(head)))

	head = array2list([]int{2, 1, 3, 5, 6, 4, 7})
	assert.Equal([]int{2, 3, 6, 7, 1, 5, 4}, list2array(oddEvenList(head)))
}
