package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}

	if r := searchRange(nums, 8); !reflect.DeepEqual(r, []int{3, 4}) {
		t.Fatal(r)
	}

	if r := searchRange(nums, 11); !reflect.DeepEqual(r, []int{-1, -1}) {
		t.Fatal(r)
	}
}
