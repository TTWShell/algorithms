package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_recoverTree(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Val: 1}
	recoverTree(root)
	assert.Equal("<<<nil> 1 <<nil> 2 <nil>>> 3 <nil>>", root.String())

	root = &TreeNode{Left: &TreeNode{Val: 1}, Val: 3, Right: &TreeNode{Left: &TreeNode{Val: 2}, Val: 4}}
	recoverTree(root)
	assert.Equal("<<<nil> 1 <nil>> 2 <<<nil> 3 <nil>> 4 <nil>>>", root.String())
}
