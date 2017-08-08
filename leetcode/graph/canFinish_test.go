package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	if r := canFinish(2, [][]int{{1, 0}}); r != true {
		t.Fatal(r)
	}

	if r := canFinish(2, [][]int{{1, 0}, {0, 1}}); r != false {
		t.Fatal(r)
	}
}
