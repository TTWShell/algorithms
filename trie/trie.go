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
	rword := []rune(word)
	if len(rword) == 0 {
		panic("Length of word must > 0.")
	}

	var tmp, curNode *trieNode
	var ok bool
	rootR := rword[0]
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
	return false

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
