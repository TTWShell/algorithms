package ldc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addOperators(t *testing.T) {
	assert := assert.New(t)

	excepted := []string{"1+2+3", "1*2*3"}
	result := addOperators("123", 6)
	assert.Subset(excepted, result)
	assert.Len(result, 2)

	excepted = []string{"2*3+2", "2+3*2"}
	result = addOperators("232", 8)
	assert.Subset(excepted, result)
	assert.Len(result, 2)

	excepted = []string{"1*0+5", "10-5"}
	result = addOperators("105", 5)
	assert.Subset(excepted, result)
	assert.Len(result, 2)

	excepted = []string{"0+0", "0-0", "0*0"}
	result = addOperators("00", 0)
	assert.Subset(excepted, result)
	assert.Len(result, 3)

	assert.Equal([]string{}, addOperators("3456237490", 9191))
}
