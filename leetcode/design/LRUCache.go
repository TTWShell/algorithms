/* https://leetcode.com/problems/lru-cache/
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present.
When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2);

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
*/

package ldesign

type LRUCache struct {
	Head *node
	Tail *node
	HT   map[int]*node
	Cap  int
}

type node struct {
	Key  int
	Val  int
	Prev *node
	Next *node
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{HT: make(map[int]*node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.HT[key]; ok {
		this.remove(node)
		this.add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.HT[key]; ok {
		node.Val = value
		this.remove(node)
		this.add(node)
		return
	}
	node := &node{Key: key, Val: value}
	this.HT[key] = node
	this.add(node)

	if len(this.HT) > this.Cap {
		delete(this.HT, this.Tail.Key)
		this.remove(this.Tail)
	}
}

func (this *LRUCache) add(node *node) {
	node.Prev = nil
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) remove(node *node) {
	if node != this.Head {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}
	if node != this.Tail {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
