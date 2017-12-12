package tree

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	bst := NewBST()

	assert.True(bst.Insert(1))
	assert.Equal(&TreeNode{Val: 1}, bst.root)

	els := []int{3, 2, 15, 18, 23, 5, 7}
	for _, el := range els {
		assert.True(bst.Insert(el))
	}

	assert.False(bst.Insert(1))

	assert.Equal([]int{1, 2, 3, 5, 7, 15, 18, 23}, InOrderRecursion(bst.root))
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	bst := NewBST()

	assert.False(bst.Search(1), "empty tree should return false")

	els := []int{3, 2, 15, 18, 23, 5, 7}
	for _, el := range els {
		assert.True(bst.Insert(el))
	}

	for _, el := range els {
		assert.True(bst.Search(el))
	}

	assert.False(bst.Search(143414), "143414 should be not exist")
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	bst := NewBST()

	assert.False(bst.Delete(1), "empty tree should return false")

	assertIsBST := func(bst *BST, el int) {
		bst.Delete(el)
		r := InOrderRecursion(bst.root)
		origin, sorted := []int{}, []int{}
		for _, num := range r {
			origin = append(origin, num)
			sorted = append(sorted, num)
		}
		sort.Ints(sorted)
		assert.Equal(sorted, origin)
	}

	els := []int{3, 2, 15, 18, 23, 5, 7, 26, 21, 1, 4, 100, 77, 87}
	for _, el := range els {
		bst.Insert(el)
	}
	for _, el := range els {
		assertIsBST(bst, el)
		bst.Insert(el)
	}
	for i := len(els) - 1; i >= 0; i-- {
		assertIsBST(bst, els[i])
		bst.Insert(els[i])
	}

	assert.False(bst.Delete(134144), "not exist el")
}
