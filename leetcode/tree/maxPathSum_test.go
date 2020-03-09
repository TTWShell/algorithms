package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Left:  &TreeNode{Val: 2},
		Val:   1,
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(6, maxPathSum(root))

	root = &TreeNode{
		Left: &TreeNode{Val: 9},
		Val:  -10,
		Right: &TreeNode{
			Left:  &TreeNode{Val: 15},
			Val:   20,
			Right: &TreeNode{Val: 7},
		},
	}
	assert.Equal(42, maxPathSum(root))

	assert.Equal(-3, maxPathSum(&TreeNode{Val: -3}))
	assert.Equal(-1, maxPathSum(&TreeNode{Val: -1, Left: &TreeNode{Val: -2}}))
	assert.Equal(1, maxPathSum(&TreeNode{Val: -2, Left: &TreeNode{Val: 1}}))
}
