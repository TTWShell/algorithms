package unionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	uf := New()
	assert.NotNil(uf.nodes)
	assert.NotNil(uf.ufSets)
	assert.Len(uf.nodes, 0)
	assert.Len(uf.ufSets, 0)
}

func TestUnion(t *testing.T) {
	assert := assert.New(t)
	uf := New()

	// set A
	uf.Union(1, 2)
	uf.Union(1, 2)
	assert.Len(uf.nodes, 2)
	assert.Len(uf.ufSets, 1)

	node1, node2 := uf.nodes[interface{}(1)], uf.nodes[interface{}(2)]
	assert.Contains(uf.ufSets, node1)
	assert.Equal(node1.value, 1)
	assert.Equal(node1.ancestor, node1)
	assert.Contains(node1.childs, node2)
	assert.NotContains(uf.ufSets, node2)
	assert.Equal(node2.value, 2)
	assert.Equal(node2.ancestor, node1)
	assert.Empty(node2.childs)

	uf.Union(1, 3)
	node3 := uf.nodes[interface{}(3)]
	assert.Len(node1.childs, 2)
	assert.Empty(node2.childs)
	assert.Len(uf.nodes, 3)
	assert.Len(uf.ufSets, 1)
	assert.Equal(node3.ancestor, node1)
	assert.Contains(node1.childs, node3)

	// set B
	//    10
	//   /  \
	// 11    12
	//      /
	//     13
	uf.Union(10, 11)
	uf.Union(10, 12)
	uf.Union(12, 13)

	assert.Len(uf.nodes, 3+4)
	node10, node11, node12, node13 := uf.nodes[interface{}(10)], uf.nodes[interface{}(11)], uf.nodes[interface{}(12)], uf.nodes[interface{}(13)]
	assert.Contains(uf.ufSets, node10)
	assert.Contains(node10.childs, node11)
	assert.Contains(node10.childs, node12)
	assert.Empty(node11.childs)
	assert.Contains(node12.childs, node13)
	assert.Empty(node13.childs)
	assert.Len(uf.ufSets, 2)
	assert.Equal(uf.ufSets[node10], 4)

	// SET C
	//        20
	//      /
	//     21
	//    /  \
	//   22   23
	//    \
	//     24
	uf.Union(20, 21)
	uf.Union(21, 22)
	uf.Union(21, 23)
	uf.Union(22, 24)

	assert.Len(uf.nodes, 3+4+5)
	node20, node21, node22, node23, node24 := uf.nodes[interface{}(20)], uf.nodes[interface{}(21)], uf.nodes[interface{}(22)], uf.nodes[interface{}(23)], uf.nodes[interface{}(24)]
	assert.Contains(uf.ufSets, node20)
	assert.Contains(node20.childs, node21)
	assert.Len(node20.childs, 1)
	assert.Contains(node21.childs, node22)
	assert.Contains(node21.childs, node23)
	assert.Len(node21.childs, 2)
	assert.Contains(node22.childs, node24)
	assert.Len(node22.childs, 1)
	assert.Empty(node23.childs)
	assert.Empty(node24.childs)
	assert.Len(uf.ufSets, 3)
	assert.Equal(uf.ufSets[node20], 5)

	// Union set B and C
	uf.Union(12, 21)
	assert.Len(uf.nodes, 3+4+5)
	assert.Len(uf.ufSets, 2)
	assert.Equal(uf.ufSets[node20], 4+5)
	assert.Len(node21.childs, 3)
	assert.Contains(node21.childs, node12)
	assert.Len(node12.childs, 2)
	assert.Contains(node12.childs, node10)
	assert.Len(node10.childs, 1)
	assert.NotContains(node10.childs, node12)
}
