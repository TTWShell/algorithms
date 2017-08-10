package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	if r := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8); !reflect.DeepEqual(r, [][]int{
		{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}) {
		t.Fatal(r)
	}
}
