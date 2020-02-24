package ldesign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BSTIterator(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Left: &TreeNode{Val: 3},
		Val:  7,
		Right: &TreeNode{
			Left:  &TreeNode{Val: 9},
			Val:   15,
			Right: &TreeNode{Val: 20}},
	}

	bst := BSTIteratorConstructor(root)
	assert.Equal(3, bst.Next())
	assert.Equal(7, bst.Next())
	assert.True(bst.HasNext())
	assert.Equal(9, bst.Next())
	assert.True(bst.HasNext())
	assert.Equal(15, bst.Next())
	assert.True(bst.HasNext())
	assert.Equal(20, bst.Next())
	assert.False(bst.HasNext())

	root = &TreeNode{
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Val:   3,
		Right: &TreeNode{Val: 4},
	}
	bst = BSTIteratorConstructor(root)
	assert.Equal(1, bst.Next())
	assert.Equal(2, bst.Next())
	assert.Equal(3, bst.Next())
	assert.Equal(4, bst.Next())
}
