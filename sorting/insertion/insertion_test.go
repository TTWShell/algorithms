package insertion

import (
	"reflect"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 2}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 2, 3, 4, 5, 6, 7, 8}) {
		t.Fatal(arr)
	}
}
