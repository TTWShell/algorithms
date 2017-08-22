package heapsort

import (
	"reflect"
	"testing"
)

func Test_HeapSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 2}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 2, 3, 4, 5, 6, 7, 8}) {
		t.Fatal(arr)
	}
}
