package lht

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	assert := assert.New(t)

	result := subdomainVisits([]string{"9001 discuss.leetcode.com"})
	sort.Strings(result)
	excepted := []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}
	sort.Strings(excepted)
	assert.Equal(excepted, result)

	result = subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"})
	sort.Strings(result)
	excepted = []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}
	sort.Strings(excepted)
	assert.Equal(excepted, result)
}
