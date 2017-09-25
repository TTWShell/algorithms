package hashtable

import (
	"fmt"
	"math"
	"sync"
)

const (
	// A = (sqrt(5) -1)/2，来自算法导论第三版
	A       = 0.6180339887
	MinSize = 10
	// Maximum average load of a bucket that triggers growth.
	LoadFactor = 0.65
)

type KeyError struct {
	s string
}

func (e *KeyError) Error() string {
	return e.s
}

type item struct {
	key   interface{}
	value interface{}
}

// 散列函数使用乘法散列法，冲突处理使用开放定址法（线性探测）的哈希表实现。
type HashTable struct {
	items []*item
	size  int
	count int
	sync.RWMutex
}

func NewHT(size int) *HashTable {
	if size >= MinSize {
		return &HashTable{size: size, items: make([]*item, size, size)}
	}
	panic(fmt.Sprintf("size must > %d", MinSize))
}

// 散列函数，支持int、string类型的key
func (ht *HashTable) hash(key interface{}) int {
	var tmp int
	switch k := key.(type) {
	case int:
		tmp = int(k)
	case string:
		// http://algs4.cs.princeton.edu/lectures/34HashTables.pdf
		// Horner's Method to hash string of length L (O(L))
		hash := int32(0)
		for i := 0; i < len(k); i++ {
			// 31 * i == (i << 5) - i
			hash = int32(k[i]) + int32(hash<<5-hash)
		}
		tmp = int(math.Abs(float64(hash)))
	default:
		panic("Only support int、string type key")
	}

	// 乘法散列法，来自算法导论第三版
	multip := float64(tmp) * A
	decimalPart := multip - float64(int(multip))
	return int(math.Floor(float64(ht.size) * decimalPart))
}

// 线性探测
func (ht *HashTable) linearProbing(hash, i int) int {
	return (hash + i) % ht.size
}

func (ht *HashTable) rehash() {
	oldSize, newSize := ht.size, ht.size<<1
	if newSize > oldSize {
		ht.size = newSize
		newItems := make([]*item, newSize, newSize)
		for _, item := range ht.items {
			if item == nil {
				continue
			}
			hash := ht.hash(item.key)
			for i := 0; i < newSize; i++ {
				tmp := ht.linearProbing(hash, i)
				if res := newItems[tmp]; res == nil {
					newItems[tmp] = item
					break
				}
			}
		}
		ht.items = newItems
	}
}

// Update if key is exists else create. Auto rehash(oldSize * 2).
func (ht *HashTable) Put(key, value interface{}) {
	ht.Lock()
	defer ht.Unlock()

	if float64(ht.count) >= float64(ht.size)*LoadFactor {
		ht.rehash()
	}

	hash := ht.hash(key)
	for i := 0; i < ht.size; i++ {
		tmp := ht.linearProbing(hash, i)
		if res := ht.items[tmp]; res == nil {
			ht.items[tmp] = &item{key: key, value: value}
			ht.count++
			return
		} else if res.key == key {
			// update
			res.value = value
			return
		}
	}
	panic("hash table overflow")
}

// Return (value, error). If key not exists, return KeyError.
func (ht *HashTable) Get(key interface{}) (value interface{}, err error) {
	ht.RLock()
	defer ht.RUnlock()

	hash := ht.hash(key)
	for i := 0; i < ht.size; i++ {
		tmp := ht.linearProbing(hash, i)
		// 如果找不到则必定不存在，否则，逐个比较key
		if ht.items[tmp] == nil {
			break
		} else if ht.items[tmp].key == key {
			return ht.items[tmp].value, nil
		}
	}
	return nil, &KeyError{"key no exists"}
}

// Return (value, error). If key not exists, return KeyError.
func (ht *HashTable) Pop(key interface{}) (value interface{}, err error) {
	ht.Lock()
	defer ht.Unlock()

	hash := ht.hash(key)
	for i := 0; i < ht.size; i++ {
		tmp := ht.linearProbing(hash, i)
		if ht.items[tmp] == nil {
			break
		} else if ht.items[tmp].key == key {
			value = ht.items[tmp].value
			ht.items[tmp] = nil
			ht.count--
			return value, nil
		}
	}
	return nil, &KeyError{"key no exists"}
}
