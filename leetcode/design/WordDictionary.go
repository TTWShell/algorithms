/* https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
Design a data structure that supports the following two operations:

	void addWord(word)
	bool search(word)

search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

For example:

	addWord("bad")
	addWord("dad")
	addWord("mad")
	search("pad") -> false
	search("bad") -> true
	search(".ad") -> true
	search("b..") -> true

Note:
    You may assume that all words are consist of lowercase letters a-z.

You should be familiar with how a Trie works. If not, please work on this problem: Implement Trie (Prefix Tree) first.
*/

package leetcode

type wdNode struct {
	Val   byte
	IsEnd bool
	Next  map[byte]*wdNode
}

// WordDictionary based on trie.
type WordDictionary struct {
	root *wdNode
}

/** Initialize your data structure here. */
func WDConstructor() WordDictionary {
	return WordDictionary{root: &wdNode{Next: make(map[byte]*wdNode)}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	var (
		rword   = []byte(word)
		curNode = this.root
		tmp     *wdNode
		ok      bool
	)

	for _, letter := range rword {
		if tmp, ok = curNode.Next[letter]; !ok {
			tmp = &wdNode{Val: letter, Next: make(map[byte]*wdNode)}
			curNode.Next[letter] = tmp
		}
		curNode = tmp
	}
	curNode.IsEnd = true
}

func (this *WordDictionary) search(curNode *wdNode, word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		rword = []byte(word)
		ok    bool
	)

	for i := 0; i < len(rword); i++ {
		letter := rword[i]
		if letter != '.' {
			if curNode, ok = curNode.Next[letter]; !ok {
				return false
			}
		} else if i < len(rword)-1 {
			for _, cur := range curNode.Next {
				if this.search(cur, string(rword[i+1:])) == true {
					return true
				}
			}
			return false
		} else if len(curNode.Next) == 0 {
			return false
		} else {
			for _, cur := range curNode.Next {
				if cur.IsEnd == true {
					return true
				}
			}
			return false
		}
	}
	return curNode.IsEnd
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
