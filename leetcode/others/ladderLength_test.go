package lothers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ladderLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	assert.Equal(5, ladderLength("qa", "sq", []string{
		"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb",
		"kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br",
		"ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr",
		"po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh",
		"sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb",
		"ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz",
		"no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an",
		"me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi",
		"am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr",
		"pa", "he", "lr", "sq", "ye"}))
	assert.Equal(0, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
