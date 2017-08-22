package quick

import (
	"reflect"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	arr := []int{4, 5, 9, 2, 6, 8, 2, 1, 7, 3}
	if QuickSort(arr); !reflect.DeepEqual(arr, []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal(arr)
	}
}
