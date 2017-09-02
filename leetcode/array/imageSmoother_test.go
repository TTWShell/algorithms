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

	input = [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}
	output = [][]int{{4, 4, 5}, {5, 6, 6}, {8, 9, 9}, {11, 12, 12}, {13, 13, 14}}
	if r := imageSmoother(input); !reflect.DeepEqual(r, output) {
		t.Fatal(r)
	}

}
