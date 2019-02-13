/* https://leetcode.com/problems/design-hashmap/description/
Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

	put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
	get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
	remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Example:

	MyHashMap hashMap = new MyHashMap();
	hashMap.put(1, 1);
	hashMap.put(2, 2);
	hashMap.get(1);            // returns 1
	hashMap.get(3);            // returns -1 (not found)
	hashMap.put(2, 1);          // update the existing value
	hashMap.get(2);            // returns 1
	hashMap.remove(2);          // remove the mapping for 2
	hashMap.get(2);            // returns -1 (not found)

Note:

	All keys and values will be in the range of [0, 1000000].
	The number of operations will be in the range of [1, 10000].
	Please do not use the built-in HashMap library.
*/

package lht

const rangeOperation = 10000

type mhmElement struct {
	key, value int
}

type MyHashMap struct {
	bucket [][]mhmElement
}

/** Initialize your data structure here. */
func HashMapConstructor() MyHashMap {
	return MyHashMap{bucket: make([][]mhmElement, rangeOperation, rangeOperation)}
}

func (this *MyHashMap) index(key int) int {
	return key % rangeOperation
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	idx := this.index(key)
	for i, e := range this.bucket[idx] {
		if e.key == key {
			this.bucket[idx][i].value = value
			return
		}
	}
	this.bucket[idx] = append(this.bucket[idx], mhmElement{key: key, value: value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	idx := this.index(key)
	for _, e := range this.bucket[idx] {
		if e.key == key {
			return e.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	idx := this.index(key)
	for i, e := range this.bucket[idx] {
		if e.key == key {
			this.bucket[idx][i].value = -1
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
