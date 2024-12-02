package lll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	// Helper function to build linked list from input
	buildList := func(input [][]interface{}) *Node {
		if len(input) == 0 {
			return nil
		}
		nodes := make([]*Node, len(input))
		for i, val := range input {
			nodes[i] = &Node{Val: val[0].(int)}
		}
		for i, val := range input {
			if i < len(input)-1 {
				nodes[i].Next = nodes[i+1]
			}
			if val[1] != nil {
				nodes[i].Random = nodes[val[1].(int)]
			}
		}
		return nodes[0]
	}

	// Helper function to convert list to [][]interface{} for comparison
	convertList := func(head *Node) [][]interface{} {
		var result [][]interface{}
		indexMap := make(map[*Node]int)
		cur := head
		index := 0
		for cur != nil {
			indexMap[cur] = index
			cur = cur.Next
			index++
		}
		cur = head
		for cur != nil {
			if cur.Random != nil {
				result = append(result, []interface{}{cur.Val, indexMap[cur.Random]})
			} else {
				result = append(result, []interface{}{cur.Val, nil})
			}

			cur = cur.Next
		}
		return result
	}

	tests := []struct {
		name     string
		input    [][]interface{}
		expected [][]interface{}
	}{
		{
			name:     "Example 1",
			input:    [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
			expected: [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			name:     "Example 2",
			input:    [][]interface{}{{1, 1}, {2, 1}},
			expected: [][]interface{}{{1, 1}, {2, 1}},
		},
		{
			name:     "Example 3",
			input:    [][]interface{}{{3, nil}, {3, 0}, {3, nil}},
			expected: [][]interface{}{{3, nil}, {3, 0}, {3, nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			copiedList := copyRandomList(head)
			assert.Equal(t, tt.expected, convertList(copiedList))
		})
	}
}
