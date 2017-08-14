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
}
