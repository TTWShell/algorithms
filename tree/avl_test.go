package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLInsert(t *testing.T) {
	assert := assert.New(t)
	avl := NewAVL()

	els := []int{130, 20, 150, 190, 230, 50, 30, 170, 180, 185, 187, 10, 5}
	expectedTree := []string{
		"<<nil> 130-0 <nil>>",
		"<<<nil> 20-0 <nil>> 130-1 <nil>>",
		"<<<nil> 20-0 <nil>> 130-1 <<nil> 150-0 <nil>>>",
		"<<<nil> 20-0 <nil>> 130-2 <<nil> 150-1 <<nil> 190-0 <nil>>>>",
		"<<<nil> 20-0 <nil>> 130-2 <<<nil> 150-0 <nil>> 190-1 <<nil> 230-0 <nil>>>>",
		"<<<nil> 20-1 <<nil> 50-0 <nil>>> 130-2 <<<nil> 150-0 <nil>> 190-1 <<nil> 230-0 <nil>>>>",
		"<<<<nil> 20-0 <nil>> 30-1 <<nil> 50-0 <nil>>> 130-2 <<<nil> 150-0 <nil>> 190-1 <<nil> 230-0 <nil>>>>",
		"<<<<nil> 20-0 <nil>> 30-1 <<nil> 50-0 <nil>>> 130-3 <<<nil> 150-1 <<nil> 170-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<nil> 20-0 <nil>> 30-1 <<nil> 50-0 <nil>>> 130-3 <<<<nil> 150-0 <nil>> 170-1 <<nil> 180-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<nil> 20-0 <nil>> 30-1 <<nil> 50-0 <nil>>> 130-3 <<<<nil> 150-0 <nil>> 170-1 <nil>> 180-2 <<<nil> 185-0 <nil>> 190-1 <<nil> 230-0 <nil>>>>>",
		"<<<<<nil> 20-0 <nil>> 30-1 <<nil> 50-0 <nil>>> 130-2 <<<nil> 150-0 <nil>> 170-1 <nil>>> 180-3 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<<<nil> 10-0 <nil>> 20-1 <nil>> 30-2 <<nil> 50-0 <nil>>> 130-3 <<<nil> 150-0 <nil>> 170-1 <nil>>> 180-4 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<<<nil> 5-0 <nil>> 10-1 <<nil> 20-0 <nil>>> 30-2 <<nil> 50-0 <nil>>> 130-3 <<<nil> 150-0 <nil>> 170-1 <nil>>> 180-4 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
	}
	for i, el := range els {
		assert.True(avl.Insert(el), avl.root.String())
		assert.Equal(avl.root.String(), expectedTree[i])
	}
	for _, el := range els {
		assert.False(avl.Insert(el), el, avl.root.String())
		assert.Equal(avl.root.String(), expectedTree[len(expectedTree)-1])
	}
}
