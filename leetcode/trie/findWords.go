/* https://leetcode.com/problems/word-search-ii/
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.

Example:

	Input:
	board = [
	['o','a','a','n'],
	['e','t','a','e'],
	['i','h','k','r'],
	['i','f','l','v']
	]
	words = ["oath","pea","eat","rain"]

	Output: ["eat","oath"]

Note:

	All inputs are consist of lowercase letters a-z.
	The values of words are distinct.
*/

package ltrie

type ByteTrie struct {
	IsRoot bool
	Root   string
	Next   map[byte]*ByteTrie
}

func ByteTrieConstructor(dict []string) ByteTrie {
	trie := &ByteTrie{IsRoot: false, Next: map[byte]*ByteTrie{}}
	for _, word := range dict {
		var (
			cur, end = trie, len(word)
			tmp      *ByteTrie
		)
		for i := range word {
			c := word[i]
			if next, ok := cur.Next[c]; ok {
				cur = next
			} else {
				tmp = &ByteTrie{IsRoot: false, Next: map[byte]*ByteTrie{}}
				cur.Next[c] = tmp
				cur = tmp
			}
		}
		cur.Root = word[0:end]
		cur.IsRoot = true
	}
	return *trie
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return []string{}
	}

	trie := ByteTrieConstructor(words)

	m, n := len(board), len(board[0])
	tmp := map[string]int{}
	var path map[[2]int]bool

	var search func(i, j int, trie *ByteTrie)
	search = func(i, j int, trie *ByteTrie) {
		char := board[i][j]
		path[[2]int{i, j}] = true
		if subTrie, ok := trie.Next[char]; ok {
			if subTrie.IsRoot {
				tmp[subTrie.Root]++
			}
			for _, point := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
				idxI, idxJ := point[0], point[1]
				if 0 <= idxI && idxI < m && 0 <= idxJ && idxJ < n && !path[[2]int{idxI, idxJ}] {
					search(idxI, idxJ, subTrie)
				}
			}
		}
		path[[2]int{i, j}] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			path = map[[2]int]bool{}
			search(i, j, &trie)
		}
	}

	res := make([]string, 0, len(tmp))
	for word := range tmp {
		res = append(res, word)
	}
	return res
}
