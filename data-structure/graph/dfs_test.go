package graph

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
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

	res := g.DFS(node1)
	assert.True(reflect.DeepEqual(res, []*Node{node1, node2, node3, node5, node4}) ||
		reflect.DeepEqual(res, []*Node{node1, node2, node5, node3, node4}) ||
		reflect.DeepEqual(res, []*Node{node1, node2, node5, node4, node3}) ||
		reflect.DeepEqual(res, []*Node{node1, node4, node2, node3, node5}) ||
		reflect.DeepEqual(res, []*Node{node1, node4, node2, node5, node3}),
	)
}
