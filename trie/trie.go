/*
FROM: https://leetcode.com/problems/implement-trie-prefix-tree/description/
*/

package trie

type trieNode struct {
	Val   rune
	IsEnd bool
	Next  map[rune]*trieNode
}

type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &trieNode{Next: make(map[rune]*trieNode)}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		panic("Length of word must > 0.")
	}

	var (
		rword   = []rune(word)
		curNode = this.root
		tmp     *trieNode
		ok      bool
	)

	for _, letter := range rword {
		if tmp, ok = curNode.Next[letter]; !ok {
			tmp = &trieNode{Val: letter, Next: make(map[rune]*trieNode)}
			curNode.Next[letter] = tmp
		}
		curNode = tmp
	}
	curNode.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		rword   = []rune(word)
		curNode = this.root
		ok      bool
	)

	for _, letter := range rword {
		if curNode, ok = curNode.Next[letter]; !ok {
			return false
		}
	}
	if !curNode.IsEnd {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	var (
		rword   = []rune(prefix)
		curNode = this.root
		ok      bool
	)

	for _, letter := range rword {
		if curNode, ok = curNode.Next[letter]; !ok {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
