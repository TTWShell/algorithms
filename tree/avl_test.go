package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLInsert(t *testing.T) {
	assert := assert.New(t)
	avl := NewAVL()

	els := []int{3, 2, 15, 18, 23, 5, 7}
	expectedTree := []string{
		"<<nil> 3-0 <nil>>",
		"<<<nil> 2-0 <nil>> 3-1 <nil>>",
		"<<<nil> 2-0 <nil>> 3-1 <<nil> 15-0 <nil>>>",
		"<<<nil> 2-0 <nil>> 3-2 <<nil> 15-1 <<nil> 18-0 <nil>>>>",
		"<<<<nil> 2-0 <nil>> 3-1 <nil>> 15-2 <<nil> 18-1 <<nil> 23-0 <nil>>>>",
		"<<<<nil> 2-0 <nil>> 3-1 <nil>> 5-3 <<nil> 15-2 <<nil> 18-1 <<nil> 23-0 <nil>>>>>",
		"<<<<nil> 2-0 <nil>> 3-1 <nil>> 5-3 <<<nil> 7-0 <nil>> 15-2 <<nil> 18-1 <<nil> 23-0 <nil>>>>>",
	}
	for i, el := range els {
		assert.True(avl.Insert(el), avl.root)
		assert.Equal(avl.root.String(), expectedTree[i])
	}
}
