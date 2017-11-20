package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeNode(t *testing.T) {
	assert := assert.New(t)

	g := New(Undirected)
	node1, node2 := g.MakeNode(1), g.MakeNode(2)
	assert.NotEqual(node1, node2)
	assert.Len(g.nodes, 2)
	assert.Len(g.edges, 2)
	assert.Len(g.edges[node1], 0)
}

func TestMakeEdge(t *testing.T) {
	assert := assert.New(t)

	g := New(Undirected)
	node1, node2 := g.MakeNode(1), g.MakeNode(2)
	assert.Len(g.edges, 2)

	node3 := &Node{3}
	assert.Error(g.MakeEdge(node3, node1, 0))
	assert.Error(g.MakeEdge(node1, node3, 0))

	assert.Nil(g.MakeEdge(node1, node2, 0.123))
	assert.Len(g.edges, 2)
	assert.Len(g.edges[node1], 1)
	assert.Len(g.edges[node2], 1, "because graph is Undirected")

	g = New(Directed)
	node1, node2 = g.MakeNode(1), g.MakeNode(2)
	assert.Nil(g.MakeEdge(node1, node2, 0.123))
	assert.Len(g.edges, 2)
	assert.Len(g.edges[node1], 1)
	assert.Len(g.edges[node2], 0, "because graph is Directed")

	assert.Nil(g.MakeEdge(node2, node1, 0.321))
	assert.Equal(g.edges[node1][node2], float32(0.123))
	assert.Equal(g.edges[node2][node1], float32(0.321))
}
