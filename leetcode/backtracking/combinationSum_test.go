package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	if r := combinationSum([]int{2, 3, 6, 7}, 7); !reflect.DeepEqual(r, [][]int{{7}, {2, 2, 3}}) {
		t.Fatal(r)
	}
}
