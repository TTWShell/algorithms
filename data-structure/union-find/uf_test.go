package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	uf.Union(1, 1)
	uf.Union(1, 2)
	uf.Union(1, 2)
	assert.Len(uf.nodes, 2)
	assert.Len(uf.ufSets, 1)

	node1, node2 := uf.nodes[interface{}(1)], uf.nodes[interface{}(2)]
	assert.Contains(uf.ufSets, node1)
	assert.Equal(node1.value, 1)
	assert.Equal(node1.ancestor, node1)
	assert.Contains(node1.children, node2)
	assert.NotContains(uf.ufSets, node2)
	assert.Equal(node2.value, 2)
	assert.Equal(node2.ancestor, node1)
	assert.Empty(node2.children)

	uf.Union(1, 3)
	node3 := uf.nodes[interface{}(3)]
	assert.Len(node1.children, 2)
	assert.Empty(node2.children)
	assert.Len(uf.nodes, 3)
	assert.Len(uf.ufSets, 1)
	assert.Equal(node3.ancestor, node1)
	assert.Contains(node1.children, node3)

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
	assert.Contains(node10.children, node11)
	assert.Contains(node10.children, node12)
	assert.Empty(node11.children)
	assert.Contains(node12.children, node13)
	assert.Empty(node13.children)
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
	assert.Contains(node20.children, node21)
	assert.Len(node20.children, 1)
	assert.Contains(node21.children, node22)
	assert.Contains(node21.children, node23)
	assert.Len(node21.children, 2)
	assert.Contains(node22.children, node24)
	assert.Len(node22.children, 1)
	assert.Empty(node23.children)
	assert.Empty(node24.children)
	assert.Len(uf.ufSets, 3)
	assert.Equal(uf.ufSets[node20], 5)

	// Union set B and C
	uf.Union(12, 21)
	assert.Len(uf.nodes, 3+4+5)
	assert.Len(uf.ufSets, 2)
	assert.Equal(uf.ufSets[node20], 4+5)
	assert.Len(node21.children, 3)
	assert.Contains(node21.children, node12)
	assert.Len(node12.children, 2)
	assert.Contains(node12.children, node10)
	assert.Len(node10.children, 1)
	assert.NotContains(node10.children, node12)
}

func TestUnion2(t *testing.T) {
	assert := assert.New(t)

	uf := New()
	// set A
	//    10
	//   /  \
	// 11    12
	//      /
	//     13
	//    /
	//   14
	uf.Union(10, 11)
	uf.Union(10, 12)
	uf.Union(12, 13)
	uf.Union(13, 14)
	node10, node11, node12, node13, node14 := uf.nodes[interface{}(10)], uf.nodes[interface{}(11)], uf.nodes[interface{}(12)], uf.nodes[interface{}(13)], uf.nodes[interface{}(14)]

	// SET B
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
	node20, node21, node22, node23, node24 := uf.nodes[interface{}(20)], uf.nodes[interface{}(21)], uf.nodes[interface{}(22)], uf.nodes[interface{}(23)], uf.nodes[interface{}(24)]

	// Union set A and B
	//        10
	//       /  \
	//     11   12
	//         /  \
	//        13   22
	//       /    /  \
	//      14   24  21
	//              /  \
	//             20  23
	uf.Union(12, 22)
	assert.Len(uf.nodes, 5+5)
	assert.Len(uf.ufSets, 1)

	assert.Len(node10.children, 2)
	assert.Equal(node10.ancestor, node10)
	assert.Equal(node10.parent, node10)
	assert.Contains(node10.children, node11)
	assert.Contains(node10.children, node12)

	assert.Empty(node11.children)
	assert.Equal(node11.ancestor, node10)
	assert.Equal(node11.parent, node10)

	assert.Len(node12.children, 2)
	assert.Equal(node12.ancestor, node10)
	assert.Equal(node12.parent, node10)
	assert.Contains(node12.children, node13)
	assert.Contains(node12.children, node22)

	assert.Len(node13.children, 1)
	assert.Equal(node13.ancestor, node10)
	assert.Equal(node13.parent, node12)
	assert.Contains(node13.children, node14)

	assert.Empty(node14.children)
	assert.Equal(node14.ancestor, node10)
	assert.Equal(node14.parent, node13)

	assert.Len(node22.children, 2)
	assert.Equal(node22.ancestor, node10)
	assert.Equal(node22.parent, node12)
	assert.Contains(node22.children, node24)
	assert.Contains(node22.children, node21)

	assert.Empty(node24.children)
	assert.Equal(node24.ancestor, node10)
	assert.Equal(node24.parent, node22)

	assert.Len(node21.children, 2)
	assert.Equal(node21.ancestor, node10)
	assert.Equal(node21.parent, node22)
	assert.Contains(node21.children, node20)
	assert.Contains(node21.children, node23)

	assert.Empty(node20.children)
	assert.Equal(node20.ancestor, node10)
	assert.Equal(node20.parent, node21)

	assert.Empty(node23.children)
	assert.Equal(node23.ancestor, node10)
	assert.Equal(node23.parent, node21)
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	uf := New()
	uf.Union(1, 2)
	assert.Equal(uf.Find(1), 1)
	assert.Equal(uf.Find(2), 1)
	assert.Nil(uf.Find(3))
}

func TestConnected(t *testing.T) {
	assert := assert.New(t)

	uf := New()
	uf.Union(1, 2)
	uf.Union(3, 4)
	assert.True(uf.Connected(1, 2))
	assert.True(uf.Connected(3, 4))
	assert.False(uf.Connected(1, 3))
	assert.False(uf.Connected(1, 4))
	assert.False(uf.Connected(2, 3))
	assert.False(uf.Connected(2, 4))
}

func TestCount(t *testing.T) {
	assert := assert.New(t)

	uf := New()
	assert.Equal(uf.Count(), 0)
	uf.Union(1, 2)
	assert.Equal(uf.Count(), 1)
	uf.Union(3, 4)
	assert.Equal(uf.Count(), 2)
	uf.Union(1, 4)
	assert.Equal(uf.Count(), 1)
	uf.Union(2, 3)
	assert.Equal(uf.Count(), 1)
}
