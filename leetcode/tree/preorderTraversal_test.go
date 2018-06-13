package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2, 3}, preorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
