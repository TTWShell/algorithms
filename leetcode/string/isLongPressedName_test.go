package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isLongPressedName(t *testing.T) {
	assert := assert.New(t)

	assert.True(isLongPressedName("alex", "aaleex"))
	assert.False(isLongPressedName("saeed", "ssaaedd"))
	assert.True(isLongPressedName("leelee", "lleeelee"))
	assert.True(isLongPressedName("laiden", "laiden"))
	assert.False(isLongPressedName("pyplrz", "ppyypllr"))
}
