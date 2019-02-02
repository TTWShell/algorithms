package lht

import "testing"
import "github.com/stretchr/testify/assert"

func Test_shortestCompletingWord(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("steps", shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	assert.Equal("pest", shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
}
