package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))
	assert.Equal([]string{"0.0.0.0"}, restoreIpAddresses("0000"))
	assert.Equal([]string{"0.10.0.10", "0.100.1.0"}, restoreIpAddresses("010010"))
}
