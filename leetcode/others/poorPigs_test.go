package lothers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_poorPigs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, poorPigs(1000, 15, 60))
	assert.Equal(0, poorPigs(1, 1, 1))
}
