package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leafSimilar(t *testing.T) {
	assert := assert.New(t)

	root1 := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 2}, Val: 3, Right: &TreeNode{Val: 4}}, Val: 5, Right: &TreeNode{Val: 6}}
	root2 := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 2}, Val: 3, Right: &TreeNode{Val: 4}}, Val: 5, Right: &TreeNode{Val: 6}}
	assert.True(leafSimilar(root1, root2))

	root3 := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 2}, Val: 3, Right: &TreeNode{Val: 4}}, Val: 5}
	assert.False(leafSimilar(root1, root3))
}
