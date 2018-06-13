package graph

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

	res := g.BFS(node1)
	assert.True(reflect.DeepEqual(res[0], node1) &&
		(reflect.DeepEqual(res[1:3], []*Node{node2, node4}) || reflect.DeepEqual(res[1:3], []*Node{node4, node2})) &&
		(reflect.DeepEqual(res[3:], []*Node{node3, node5}) || reflect.DeepEqual(res[3:], []*Node{node5, node3})))

	res = g.BFS(node2)
	assert.True(reflect.DeepEqual(res[0], node2) &&
		(reflect.DeepEqual(res[1:3], []*Node{node3, node5}) || reflect.DeepEqual(res[1:3], []*Node{node5, node3})) &&
		(reflect.DeepEqual(res[3:], []*Node{node1, node4}) || reflect.DeepEqual(res[3:], []*Node{node4, node1})))

	res = g.BFS(node3)
	assert.True(reflect.DeepEqual(res[0:2], []*Node{node3, node1}) &&
		(reflect.DeepEqual(res[2:4], []*Node{node2, node4}) || reflect.DeepEqual(res[2:4], []*Node{node4, node2})) &&
		(reflect.DeepEqual(res[4:], []*Node{node5}) || reflect.DeepEqual(res[4:], []*Node{node5})))

	assert.Equal(g.BFS(node4), []*Node{node4})

	res = g.BFS(node5)
	assert.True(reflect.DeepEqual(res[0], node5) &&
		(reflect.DeepEqual(res[1:3], []*Node{node3, node4}) || reflect.DeepEqual(res[1:3], []*Node{node4, node3})) &&
		reflect.DeepEqual(res[3:], []*Node{node1, node2}))
}
