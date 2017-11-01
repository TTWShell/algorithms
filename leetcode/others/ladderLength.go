/* https://leetcode.com/problems/word-ladder/description/
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
For example,

    Given:
    beginWord = "hit"
    endWord = "cog"
    wordList = ["hot","dot","dog","lot","log","cog"]
    As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
    return its length 5.

Note:
    Return 0 if there is no such transformation sequence.
    All words have the same length.
    All words contain only lowercase alphabetic characters.
    You may assume no duplicates in the word list.
    You may assume beginWord and endWord are non-empty and are not the same.
UPDATE (2017/1/20):
    The wordList parameter had been changed to a list of strings (instead of a set of strings). Please reload the code definition to get the latest changes.
*/

package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		words[word] = true
	}

	var (
		wordTmp    string
		res, queue = 0, []string{beginWord}
	)

	for len(queue) != 0 {
		res++
		newQueue := []string{}
		for _, word := range queue {
			copyList := []byte(word)
			for i, cur := range copyList {
				for j := 'a'; j <= 'z'; j++ {
					tmp := byte(j)
					if tmp == cur {
						continue
					}
					copyList[i], wordTmp = tmp, string(copyList)
					if _, ok := words[wordTmp]; ok {
						if wordTmp == endWord {
							return res + 1
						}
						newQueue = append(newQueue, wordTmp)
						delete(words, wordTmp)
					}
				}
				copyList[i] = cur
			}
		}
		queue = newQueue
	}
	return 0
}
