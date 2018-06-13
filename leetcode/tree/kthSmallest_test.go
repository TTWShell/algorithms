package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	assert := assert.New(t)

	/*
	         4
	       /   \
	     2      5
	    / \
	   1   3
	*/
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5}}
	for i := 1; i <= 5; i++ {
		assert.Equal(i, kthSmallest(root, i))
	}

	assert.Panicsf(func() { kthSmallest(root, 100) }, "error occurred")
}
