package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	if nextPermutation(nums); !reflect.DeepEqual(nums, []int{1, 3, 2}) {
		t.Fatal(nums)
	}

	nums = []int{3, 2, 1}
	if nextPermutation(nums); !reflect.DeepEqual(nums, []int{1, 2, 3}) {
		t.Fatal(nums)
	}

	nums = []int{1, 1, 5}
	if nextPermutation(nums); !reflect.DeepEqual(nums, []int{1, 5, 1}) {
		t.Fatal(nums)
	}
}
