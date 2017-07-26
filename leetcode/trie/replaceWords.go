/* https://leetcode.com/problems/replace-words/#/description
In English, we have a concept called root, which can be followed by some other words to form another longer word - let's call this word successor.
For example, the root an, followed by other, which can form another word another.

Now, given a dictionary consisting of many roots and a sentence. You need to replace all the successor in the sentence with the root forming it.
If a successor has many roots can form it, replace it with the root with the shortest length.

You need to output the sentence after the replacement.

Example 1:
    Input: dict = ["cat", "bat", "rat"]
    sentence = "the cattle was rattled by the battery"
    Output: "the cat was rat by the bat"
Note:
    The input will only have lower-case letters.
    1 <= dict words number <= 1000
    1 <= sentence words number <= 1000
    1 <= root length <= 100
    1 <= sentence words length <= 1000
*/

package leetcode

import "strings"

type Trie struct {
	IsRoot bool
	Root   string
	Next   map[rune]*Trie
}

func TrieConstructor(dict []string) Trie {
	trie := &Trie{IsRoot: false, Next: map[rune]*Trie{}}
	for _, word := range dict {
		var (
			cur, end = trie, len(word)
			tmp      *Trie
		)
		for i, c := range word {
			if next, ok := cur.Next[c]; ok {
				if cur = next; cur.IsRoot {
					// only need shortest
					end = i + 1
					break
				}
			} else {
				tmp = &Trie{IsRoot: false, Next: map[rune]*Trie{}}
				cur.Next[c] = tmp
				cur = tmp
			}
		}
		cur.Root = word[0:end]
		cur.IsRoot = true
	}
	return *trie
}

func replaceWords(dict []string, sentence string) string {
	if len(sentence) == 0 {
		return sentence
	}

	trie := TrieConstructor(dict)
	words := strings.Split(sentence, " ")

	for i, word := range words {
		cur := &trie
		for _, c := range word {
			if next, ok := cur.Next[c]; ok {
				if next.IsRoot {
					words[i] = next.Root
					break
				}
			} else {
				break
			}
			cur = cur.Next[c]
		}
	}
	return strings.Join(words, " ")
}

/*
func replaceWords(dict []string, sentence string) string {
	replace := func(root string, words []string) []string {
		if len(root) == 0 {
			return words
		}
		for i := range words {
			if strings.HasPrefix(words[i], root) {
				words[i] = root
			}
		}
		return words
	}
	words := strings.Split(sentence, " ")
	for _, root := range dict {
		words = replace(root, words)
	}
	return strings.Join(words, " ")
}
*/
