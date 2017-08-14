package leetcode

import "testing"

func Test_judgeCircle(t *testing.T) {
	if r := judgeCircle("UD"); r != true {
		t.Fatal(r)
	}

	if r := judgeCircle("LL"); r != false {
		t.Fatal(r)
	}
}
