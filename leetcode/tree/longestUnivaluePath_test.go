package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestUnivaluePath(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 1}, Val: 4, Right: &TreeNode{Val: 1}}, Val: 5, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}}
	assert.Equal(2, longestUnivaluePath(root))

	root = &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 4}, Val: 4, Right: &TreeNode{Val: 4}}, Val: 1, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}}
	assert.Equal(2, longestUnivaluePath(root))
}
