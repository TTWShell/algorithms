package ldesign

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomizedSet(t *testing.T) {
	assert := assert.New(t)

	obj := RandomizedSetConstructor()

	assert.True(obj.Insert(1))
	assert.False(obj.Remove(2))
	assert.True(obj.Insert(2))
	assert.True(obj.GetRandom() == 1 || obj.GetRandom() == 2)
	assert.True(obj.Remove(1))
	assert.False(obj.Insert(2))
	// Since 2 is the only number in the set, getRandom always return 2.
	fmt.Println(obj)
	assert.Equal(2, obj.GetRandom())
	// for i := 0; i < 10; i++ {
	// 	assert.Equal(2, obj.GetRandom())
	// }
}
