package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFS(t *testing.T) {
	assert := assert.New(t)

	g := New(Directed)
	node1, node2, node3, node4, node5 := g.MakeNode(1), g.MakeNode(2), g.MakeNode(3), g.MakeNode(4), g.MakeNode(5)
	g.MakeEdge(node1, node2, 0)
	g.MakeEdge(node1, node4, 0)
	g.MakeEdge(node2, node3, 0)
	g.MakeEdge(node2, node5, 0)
	g.MakeEdge(node3, node1, 0)
	g.MakeEdge(node5, node3, 0)
	g.MakeEdge(node5, node4, 0)

	assert.Equal(g.BFS(node1), []*Node{node1, node2, node4, node3, node5})
	assert.Equal(g.BFS(node2), []*Node{node2, node3, node5, node1, node4})
	assert.Equal(g.BFS(node3), []*Node{node3, node1, node2, node4, node5})
	assert.Equal(g.BFS(node4), []*Node{node4})
	assert.Equal(g.BFS(node5), []*Node{node5, node3, node4, node1, node2})
}
