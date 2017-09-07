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

	avl.root = nil
	// test sorted
	for el := 1; el < 30; el++ {
		avl.Insert(el)
	}
	assert.Equal(avl.root.String(), "<<<<<<nil> 1-0 <nil>> 2-1 <<nil> 3-0 <nil>>> 4-2 <<<nil> 5-0 <nil>> 6-1 <<nil> 7-0 <nil>>>>"+
		" 8-3 <<<<nil> 9-0 <nil>> 10-1 <<nil> 11-0 <nil>>> 12-2 <<<nil> 13-0 <nil>> 14-1 <<nil> 15-0 <nil>>>>>"+
		" 16-4 <<<<<nil> 17-0 <nil>> 18-1 <<nil> 19-0 <nil>>> 20-2 <<<nil> 21-0 <nil>> 22-1 <<nil> 23-0 <nil>>>>"+
		" 24-3 <<<nil> 25-0 <nil>> 26-2 <<<nil> 27-0 <nil>> 28-1 <<nil> 29-0 <nil>>>>>>")

}

func TestAVLSearch(t *testing.T) {
	assert := assert.New(t)
	avl := NewAVL()
	assert.False(avl.Search(0), 0, avl.root)

	els := []int{130, 20, 150, 190, 230, 50, 30, 170, 180, 185, 187, 10, 5}
	for _, el := range els {
		avl.Insert(el)
	}

	for _, el := range els {
		assert.True(avl.Search(el), avl.root.String())
	}

	for _, el := range []int{0, 100000} {
		assert.False(avl.Search(el), avl.root.String())
	}
}

func TestAVLDelete(t *testing.T) {
	assert := assert.New(t)
	avl := NewAVL()
	assert.False(avl.Delete(0), 0, avl.root)

	els := []int{130, 20, 150, 190, 230, 50, 30, 170, 180, 185, 187, 10, 5, 25}
	for _, el := range els {
		avl.Insert(el)
	}

	expectedTree := []string{
		"<<<<<<nil> 5-0 <nil>> 10-1 <nil>> 20-2 <<<nil> 25-0 <nil>> 30-1 <nil>>> 50-3 <<<nil> 150-0 <nil>> 170-1 <nil>>> 180-4 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<nil> 30-0 <nil>>> 50-3 <<<nil> 150-0 <nil>> 170-1 <nil>>> 180-4 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<<nil> 30-0 <nil>> 50-1 <<nil> 170-0 <nil>>>> 180-3 <<<nil> 185-1 <<nil> 187-0 <nil>>> 190-2 <<nil> 230-0 <nil>>>>",
		"<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<<nil> 30-0 <nil>> 50-1 <<nil> 170-0 <nil>>>> 180-3 <<<nil> 185-0 <nil>> 187-1 <<nil> 230-0 <nil>>>>",
		"<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<<nil> 30-0 <nil>> 50-1 <<nil> 170-0 <nil>>>> 180-3 <<<nil> 185-0 <nil>> 187-1 <nil>>>",
		"<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<<nil> 30-0 <nil>> 170-1 <nil>>> 180-3 <<<nil> 185-0 <nil>> 187-1 <nil>>>",
		"<<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<nil> 170-0 <nil>>> 180-3 <<<nil> 185-0 <nil>> 187-1 <nil>>>",
		"<<<<nil> 5-0 <nil>> 10-1 <<nil> 25-0 <nil>>> 180-2 <<<nil> 185-0 <nil>> 187-1 <nil>>>",
		"<<<<nil> 5-0 <nil>> 10-1 <<nil> 25-0 <nil>>> 185-2 <<nil> 187-0 <nil>>>",
		"<<<<nil> 5-0 <nil>> 10-1 <nil>> 25-2 <<nil> 187-0 <nil>>>",
		"<<<nil> 5-0 <nil>> 10-1 <<nil> 25-0 <nil>>>",
		"<<<nil> 5-0 <nil>> 25-1 <nil>>",
		"<<nil> 25-0 <nil>>",
		"<nil>",
	}
	for i := 0; i < len(els)-1; i++ {
		assert.True(avl.Delete(els[i]))
		assert.Equal(avl.root.String(), expectedTree[i])
	}
	assert.True(avl.Delete(25))
	assert.False(avl.Search(25))

	for _, el := range els {
		avl.Insert(el)
	}
	for _, el := range els {
		assert.True(avl.Delete(el))
		assert.True(avl.Insert(el))
	}

	avl.root = nil
	for el := 1; el < 30; el++ {
		avl.Insert(el * 10)
	}
	assert.True(avl.Delete(270))
	assert.True(avl.Delete(260)) // for test swap right

	assert.True(avl.Delete(210))
	assert.True(avl.Delete(190))
	assert.True(avl.Delete(230))
	assert.True(avl.Delete(200)) // for test swap left

	avl.root = nil
	for el := 1; el < 30; el++ {
		avl.Insert(el * 10)
	}
	assert.True(avl.Delete(210))
	assert.True(avl.Delete(230))
	assert.True(avl.Delete(220)) // for test leftRigthRotate(root)

	assert.True(avl.Delete(250)) // for test rightLeftRotate(root)

	for _, el := range []int{0, 100000} {
		assert.False(avl.Delete(el))
	}
}
