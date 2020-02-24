package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fullJustify(t *testing.T) {
	assert := assert.New(t)

	excepted := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	assert.Equal(excepted, fullJustify([]string{
		"This", "is", "an", "example", "of", "text", "justification.",
	}, 16))

	excepted = []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	assert.Equal(excepted, fullJustify([]string{
		"What", "must", "be", "acknowledgment", "shall", "be",
	}, 16))

	excepted = []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}
	assert.Equal(excepted, fullJustify([]string{
		"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
	}, 20))
}
