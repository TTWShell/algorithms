/* https://leetcode.com/problems/word-ladder-ii/
Given two words (beginWord and endWord), and a dictionary's word list,
find all shortest transformation sequence(s) from beginWord to endWord,
such that:

	Only one letter can be changed at a time
	Each transformed word must exist in the word list.
	Note that beginWord is not a transformed word.

Note:

	Return an empty list if there is no such transformation sequence.
	All words have the same length.
	All words contain only lowercase alphabetic characters.
	You may assume no duplicates in the word list.
	You may assume beginWord and endWord are non-empty and are not the same.

Example 1:

	Input:
	beginWord = "hit",
	endWord = "cog",
	wordList = ["hot","dot","dog","lot","log","cog"]

	Output:
	[
	["hit","hot","dot","dog","cog"],
	["hit","hot","lot","log","cog"]
	]

Example 2:

	Input:
	beginWord = "hit"
	endWord = "cog"
	wordList = ["hot","dot","dog","lot","log"]

	Output: []

	Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/

package lbacktracking

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	canTrans := func(word1, word2 string) bool {
		// word1和word2是否可以相互转换
		cnt := 0
		for i := 0; i < len(word1) && cnt < 2; i++ {
			if word1[i] != word2[i] {
				cnt++
			}
		}
		return cnt == 1
	}

	// 处理beginWord，endWord不在队列的情况
	beginWordIdx, endWordEIdx := -1, -1
	for i, word := range wordList {
		if word == endWord {
			endWordEIdx = i
		} else if word == beginWord {
			beginWordIdx = i
		}
	}
	if endWordEIdx == -1 {
		return [][]string{}
	}
	if beginWordIdx == -1 {
		beginWordIdx = len(wordList)
		wordList = append(wordList, beginWord)
	}

	// 缓存每个单词的候选词
	relations := make(map[int]map[int]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		relations[i] = map[int]bool{}
	}
	for i, word := range wordList {
		for j := i + 1; j < len(wordList); j++ {
			if canTrans(word, wordList[j]) {
				relations[i][j], relations[j][i] = true, true
			}
		}
	}

	type Node struct {
		value    string
		height   int
		preNodes map[*Node]bool
	}

	allNodes := make(map[int]*Node, len(wordList)+1)
	for i, word := range wordList {
		allNodes[i] = &Node{value: word, height: len(wordList), preNodes: map[*Node]bool{}}
	}
	allNodes[beginWordIdx].height = 0

	// bfs层序处理关系
	worked, curWords := map[*Node]bool{}, []int{beginWordIdx}
	for len(curWords) != 0 {
		nextWords := []int{}
		for _, curWord := range curWords {
			curNode := allNodes[curWord]
			if curNode.value == endWord {
				break
			}
			worked[curNode] = true
			for nextWord := range relations[curWord] {
				nextNode := allNodes[nextWord]
				if _, ok := worked[nextNode]; ok {
					continue
				}
				nextWords = append(nextWords, nextWord)
				if tmp := curNode.height + 1; nextNode.height >= tmp {
					nextNode.height = tmp
					nextNode.preNodes[curNode] = true
				}
			}
		}
		curWords = nextWords
	}

	var dfs func(node *Node) [][]string
	dfs = func(node *Node) [][]string {
		if node == allNodes[beginWordIdx] {
			return [][]string{{beginWord}}
		}

		res := [][]string{}
		for nextNode := range node.preNodes {
			if nextNode.height == node.height-1 {
				subs := dfs(nextNode)
				for _, sub := range subs {
					tmp := append([]string{node.value}, sub...)
					res = append(res, tmp)
				}
			}
		}
		return res
	}

	res := dfs(allNodes[endWordEIdx])
	for i := range res {
		for j := 0; j < len(res[i])/2; j++ {
			res[i][j], res[i][len(res[i])-1-j] = res[i][len(res[i])-1-j], res[i][j]
		}
	}
	return res
}
