package radix

import (
	"reflect"
	"testing"
)

func Test_RadixSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 2, 126, 35, 3, 10, 856, 72, 21, 4, 2342}
	RadixSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 10, 21, 35, 72, 126, 856, 2342}) {
		t.Fatal(arr)
	}
}
