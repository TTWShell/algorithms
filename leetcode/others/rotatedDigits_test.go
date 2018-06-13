package lothers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotatedDigits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, rotatedDigits(10))
}
