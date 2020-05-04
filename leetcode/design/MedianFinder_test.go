package ldesign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	assert := assert.New(t)

	obj := MedianFinderConstructor()

	obj.AddNum(1)
	obj.AddNum(2)
	assert.Equal(1.5, obj.FindMedian())

	obj.AddNum(3)
	assert.Equal(2.0, obj.FindMedian())
}
