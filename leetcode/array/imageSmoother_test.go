package leetcode

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	input := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	output := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	if r := imageSmoother(input); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}
}
