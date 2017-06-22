package leetcode

import (
	"sort"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	num := 1
	result := []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}
	sort.Strings(result)

	r := readBinaryWatch(num)

	if len(r) != len(result) {
		t.Fatal(num, r)
	}

	for i := range result {
		if r[i] != result[i] {
			t.Fatal(num, r)
		}
	}

	for num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		t.Log(num)
		readBinaryWatch(num)
	}
}
