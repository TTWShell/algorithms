package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Int int

func (i Int) LessThan(e Element) bool {
	return i < e.(Int)
}

func Test_Init(t *testing.T) {
	assert := assert.New(t)

	h := New(true)
	assert.False(h.len != 0 || h.Len() != 0 || len(h.heap) != 0 || h.isMin != true || !h.IsEmpty(), "init MinHeap failed.", h)

	h = New(false)
	assert.False(h.len != 0 || h.Len() != 0 || len(h.heap) != 0 || h.isMin != false || !h.IsEmpty(), "init MaxHeap failed.", h)
}

func Test_MinHeap(t *testing.T) {
	assert := assert.New(t)

	h := New(true)
	for _, num := range []Int{100, 90, 80, 5, 3, 1, 100, 70, 8, 80, 10, 20, 30, 8, 70} {
		h.Insert(num)
	}

	res := make([]int, 0, h.Len())
	for h.Len() > 0 {
		res = append(res, int(h.Extract().(Int)))
	}

	for i := 0; i < len(res)-1; i++ {
		assert.True(res[i] <= res[i+1], "error", res)
	}
}

func Test_MaxHeap(t *testing.T) {
	assert := assert.New(t)

	h := New(false)
	for _, num := range []Int{100, 90, 80, 5, 3, 1, 100, 70, 8, 80, 10, 20, 30, 8, 70} {
		h.Insert(num)
	}

	res := make([]Int, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract().(Int))
	}

	for i := 0; i < len(res)-1; i++ {
		assert.True(res[i] >= res[i+1], "error", res)
	}
}

func Test_ExtractPanic(t *testing.T) {
	assert := assert.New(t)

	h := New(true)
	defer func() {
		assert.Equal("Empty heap, cannot Extract.", recover())
	}()
	h.Extract()
}

func Test_Peek(t *testing.T) {
	assert := assert.New(t)

	h := New(true)
	h.Insert(Int(100))

	assert.Equal(Int(100), h.Peek().(Int))

	h.Extract()

	defer func() {
		assert.Equal("Empty heap, cannot Peek.", recover())
	}()
	h.Peek()
}

func Test_precolate(t *testing.T) {
	assert := assert.New(t)

	h := New(true)
	for i, num := range []Int{100, 90, 80, 5, 3, 1, 10, 20, 30, 8, 70} {
		h.Insert(num)
		h.precolateUp(i)
		h.precolateDown(h.Len() - 1 - i)
	}

	res := make([]Int, 0)
	for h.Len() > 0 {
		res = append(res, h.Extract().(Int))
	}

	for i := 0; i < len(res)-1; i++ {
		assert.True(res[i] < res[i+1])
	}
}
