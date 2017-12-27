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
	data map[rune]*trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{data: make(map[rune]*trieNode)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		panic("Length of word must > 0.")
	}

	var (
		rword        = []rune(word)
		tmp, curNode *trieNode
		ok           bool
		rootR        = rword[0]
	)
	if curNode, ok = this.data[rootR]; !ok {
		curNode = &trieNode{Val: rootR, Next: make(map[rune]*trieNode)}
		this.data[rootR] = curNode
	}
	tmp = curNode

	for i := 1; i < len(rword); i++ {
		cur := rword[i]
		if curNode, ok = tmp.Next[cur]; !ok {
			curNode = &trieNode{Val: cur, Next: make(map[rune]*trieNode)}
			tmp.Next[cur] = curNode
		}
		tmp = curNode
	}
	tmp.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		rword   = []rune(word)
		rootR   = rword[0]
		curNode *trieNode
		ok      bool
	)
	if curNode, ok = this.data[rootR]; !ok {
		return false
	}
	for i := 1; i < len(rword); i++ {
		cur := rword[i]
		if curNode, ok = curNode.Next[cur]; !ok {
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
		rootR   = rword[0]
		curNode *trieNode
		ok      bool
	)
	if curNode, ok = this.data[rootR]; !ok {
		return false
	}
	for i := 1; i < len(rword); i++ {
		cur := rword[i]
		if curNode, ok = curNode.Next[cur]; !ok {
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
