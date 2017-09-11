package hashtable

import (
	"fmt"
	"math"
)

const (
	A       = 0.6180339887 // A = (sqrt(t) -1)/2，来自算法导论第三版
	MinSize = 10
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

type HashTable struct {
	items []*item
	size  int
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
	}

	// 乘法散列法，来自算法导论第三版
	multip := float64(tmp) * A
	decimalPart := multip - float64(int(multip))
	return int(math.Floor(float64(ht.size) * decimalPart))
}

func (ht *HashTable) linearProbing(hash, i int) int {
	return (hash + i) % ht.size
}

func (ht *HashTable) Put(key, value interface{}) {
	hash := ht.hash(key)
	if ht.items[hash] == nil {
		ht.items[hash] = &item{key: key, value: value}
	}
}

func (ht *HashTable) Get(key interface{}) (value interface{}, err error) {
	hash := ht.hash(key)
	if ht.items[hash] != nil && ht.items[hash].key == key {
		return ht.items[hash].value, nil
	}
	return nil, &KeyError{"key no exists"}
}
