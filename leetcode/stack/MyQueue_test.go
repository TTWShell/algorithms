package lstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyQueue(t *testing.T) {
	assert := assert.New(t)

	myqueue := MyQueueConstructor()
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)

	assert.False(myqueue.Empty())
	assert.Equal(1, myqueue.Peek())
	assert.Equal(1, myqueue.Pop())
	assert.Equal(2, myqueue.Pop())
	assert.Equal(3, myqueue.Pop())
	assert.True(myqueue.Empty())
}
