package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]*TreeNode{&TreeNode{Val: 1}}, generateTrees(1))

	assert.Equal([]*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		&TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
	}, generateTrees(2))

	excepted := []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
		&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
	}
	assert.Equal(excepted, generateTrees(3))
}
