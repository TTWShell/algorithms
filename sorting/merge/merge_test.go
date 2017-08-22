package merge

import (
	"reflect"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	arr1 := []int{1, 3, 5, 6, 8, 10, 12, 13}
	arr2 := []int{2, 4, 5, 7, 9, 11, 14, 15, 16, 17, 18}
	output := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	if r := MergeSort(arr1, arr2); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}

	if r := MergeSort(arr2, arr1); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}
}
