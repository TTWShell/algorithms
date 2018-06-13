package lstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calPoints(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(calPoints([]string{"5", "2", "C", "D", "+"}), 30)
	assert.Equal(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}), 27)
}
