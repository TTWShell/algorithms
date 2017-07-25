package leetcode

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	if r := findErrorNums([]int{1, 2, 2, 4}); !reflect.DeepEqual(r, []int{2, 3}) {
		t.Fatal(r)
	}
}
