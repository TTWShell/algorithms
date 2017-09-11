package hashtable

import (
	"fmt"
	"math"
)

const (
	A       = 0.6180339887 // A = (sqrt(t) -1)/2，来自算法导论第三版
	MinSize = 10
)

type item struct {
	key   interface{}
	value interface{}
}

type HashTable struct {
	size int
}

func NewHT(size int) *HashTable {
	if size > MinSize {
		return &HashTable{size: size}
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
