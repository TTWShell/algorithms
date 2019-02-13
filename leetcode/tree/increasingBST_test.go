package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_increasingBST(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{Val: 2}, Val: 3, Right: &TreeNode{Val: 4}}, Val: 5, Right: &TreeNode{Val: 6}}
	excepted := &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}
	assert.Equal(excepted, increasingBST(root))
}
