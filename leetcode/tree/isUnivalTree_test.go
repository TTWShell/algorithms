package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isUnivalTree(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 1}, Val: 1, Right: &TreeNode{Val: 1}}, Val: 1, Right: &TreeNode{Left: &TreeNode{Val: 1}, Val: 1}}
	assert.True(isUnivalTree(root))

	root = &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 5}, Val: 2, Right: &TreeNode{Val: 2}}, Val: 2, Right: &TreeNode{Val: 2}}
	assert.False(isUnivalTree(root))
}
