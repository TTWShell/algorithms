package ltrie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWords(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{}, findWords([][]byte{}, []string{"cdba"}))

	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	excepted := []string{"eat", "oath"}
	result := findWords(board, []string{"oath", "pea", "eat", "rain"})
	assert.Subset(excepted, result)
	assert.Len(result, 2)

	board = [][]byte{{'a', 'a'}}
	assert.Equal([]string{}, findWords(board, []string{"aaa"}))

	board = [][]byte{{'a', 'b'}, {'c', 'd'}}
	assert.Equal([]string{"cdba"}, findWords(board, []string{"cdba"}))

	board = [][]byte{
		{'b', 'a', 'a', 'b', 'a', 'b'},
		{'a', 'b', 'a', 'a', 'a', 'a'},
		{'a', 'b', 'a', 'a', 'a', 'b'},
		{'a', 'b', 'a', 'b', 'b', 'a'},
		{'a', 'a', 'b', 'b', 'a', 'b'},
		{'a', 'a', 'b', 'b', 'b', 'a'},
		{'a', 'a', 'b', 'a', 'a', 'b'},
	}
	words := []string{
		"aab", "bbaabaabaaaaabaababaaaaababb", "aabbaaabaaabaabaaaaaabbaaaba",
		"babaababbbbbbbaabaababaabaaa", "bbbaaabaabbaaababababbbbbaaa",
		"babbabbbbaabbabaaaaaabbbaaab", "bbbababbbbbbbababbabbbbbabaa",
		"babababbababaabbbbabbbbabbba", "abbbbbbaabaaabaaababaabbabba",
		"aabaabababbbbbbababbbababbaa", "aabbbbabbaababaaaabababbaaba",
		"ababaababaaabbabbaabbaabbaba", "abaabbbaaaaababbbaaaaabbbaab",
		"aabbabaabaabbabababaaabbbaab", "baaabaaaabbabaaabaabababaaaa",
		"aaabbabaaaababbabbaabbaabbaa", "aaabaaaaabaabbabaabbbbaabaaa",
		"abbaabbaaaabbaababababbaabbb", "baabaababbbbaaaabaaabbababbb",
		"aabaababbaababbaaabaabababab", "abbaaabbaabaabaabbbbaabbbbbb",
		"aaababaabbaaabbbaaabbabbabab", "bbababbbabbbbabbbbabbbbbabaa",
		"abbbaabbbaaababbbababbababba", "bbbbbbbabbbababbabaabababaab",
		"aaaababaabbbbabaaaaabaaaaabb", "bbaaabbbbabbaaabbaabbabbaaba",
		"aabaabbbbaabaabbabaabababaaa", "abbababbbaababaabbababababbb",
		"aabbbabbaaaababbbbabbababbbb", "babbbaabababbbbbbbbbaabbabaa",
	}
	result = findWords(board, words)
	excepted = []string{"aab", "aabbbbabbaababaaaabababbaaba", "abaabbbaaaaababbbaaaaabbbaab", "ababaababaaabbabbaabbaabbaba"}
	assert.Subset(excepted, result)
	assert.Len(result, 4)
}
