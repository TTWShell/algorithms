package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeNode(t *testing.T) {
	assert := assert.New(t)

	g := New()
	node1, node2 := g.MakeNode(1), g.MakeNode(2)
	assert.NotEqual(node1, node2)
	assert.Len(g.nodes, 2)
	assert.Len(g.edges, 2)
	assert.Len(g.edges[node1], 0)
}
