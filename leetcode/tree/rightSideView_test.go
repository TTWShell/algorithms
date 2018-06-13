package ltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Val:   1,
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	assert.Equal([]int{1, 3, 4}, rightSideView(root))

	assert.Equal([]int{}, rightSideView(nil))
}
