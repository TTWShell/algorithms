package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numUniqueEmails(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	assert.Equal(3, numUniqueEmails([]string{"testemail@leetcode.com", "testemail1@leetcode.com", "testemail+david@lee.tcode.com"}))
}
