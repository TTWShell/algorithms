package leetcode

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	if r := findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3); !reflect.DeepEqual(r, []int{1, 2, 3, 4}) {
		t.Fatal(r)
	}

	if r := findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1); !reflect.DeepEqual(r, []int{1, 2, 3, 4}) {
		t.Fatal(r)
	}

	if r := findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5); !reflect.DeepEqual(r, []int{3, 3, 4}) {
		t.Fatal(r)
	}

	if r := findClosestElements([]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2); !reflect.DeepEqual(r, []int{1, 3}) {
		t.Fatal(r)
	}
}
