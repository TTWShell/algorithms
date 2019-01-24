package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList(t *testing.T) {
	assert := assert.New(t)

	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	assert.Equal(2, linkedList.Get(1)) // linked list becomes 1->2->3
	linkedList.DeleteAtIndex(1)        // now the linked list is 1->3
	assert.Equal(3, linkedList.Get(1))

	linkedList.AddAtIndex(1, 2)
	linkedList.DeleteAtIndex(2)
	assert.Equal(-1, linkedList.Get(2))

	linkedList.AddAtIndex(1, 2)
	linkedList.DeleteAtIndex(0)
	assert.Equal(2, linkedList.Get(0))
}
