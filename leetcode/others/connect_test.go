package lothers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected []interface{}
	}{
		{
			name: "Example 1",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},
			},
			expected: []interface{}{1, "#", 2, 3, "#", 4, 5, 7, "#"},
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := connect(tt.root)

			// Serialize the tree and check the result
			result := serializeTree(root)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func serializeTree(root *Node) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	queue := []*Node{root}
	var result []interface{}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			result = append(result, node.Val)
			if i == levelSize-1 {
				result = append(result, "#")
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelSize:]
	}
	return result
}
