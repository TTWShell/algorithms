package merge

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	assert := assert.New(t)

	arr := []int{3, 13, 13, 13, 0, 2, 7, 10, 10, 11, 14, 2, 2, 12, 6, 9}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	assert.Equal([]int{3, 13, 13, 13, 0, 2, 7, 10, 10, 11, 14, 2, 2, 12, 6, 9}, arr)
	assert.Equal([]int{0, 2, 2, 2, 3, 6, 7, 9, 10, 10, 11, 12, 13, 13, 13, 14}, sorted)
	assert.Equal(sorted, Sort(arr))
}
