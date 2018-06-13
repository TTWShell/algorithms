package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getHint(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0A4B", getHint("1122", "2211"))
	assert.Equal("3A0B", getHint("1122", "1222"))
	assert.Equal("1A3B", getHint("1807", "7810"))
	assert.Equal("1A1B", getHint("1123", "0111"))

	assert.Equal("2A12B", getHint("50135916782508955012", "76298008752757713977"))

}
