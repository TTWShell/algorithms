package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberToWords(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Zero", numberToWords(0))
	assert.Equal("One Hundred Twenty Three", numberToWords(123))
	assert.Equal("Twelve Thousand Three Hundred Forty Five", numberToWords(12345))
	assert.Equal("One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven", numberToWords(1234567))
	assert.Equal("One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One", numberToWords(1234567891))
}
