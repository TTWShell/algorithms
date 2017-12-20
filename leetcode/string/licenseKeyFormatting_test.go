package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_licenseKeyFormatting(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("5F3Z-2E9W", licenseKeyFormatting("5F3Z-2e-9-w", 4))
	assert.Equal("2-5G-3J", licenseKeyFormatting("2-5g-3-J", 2))
	assert.Equal("R", licenseKeyFormatting("r", 1))
	assert.Equal("AA-AA", licenseKeyFormatting("aaaa", 2))
	assert.Equal("", licenseKeyFormatting("---", 1))
}
