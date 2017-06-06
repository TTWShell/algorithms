package leetcode

import "testing"

func Test_findWords(t *testing.T) {
	r := findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	if len(r) != 2 || r[0] != "Alaska" || r[1] != "Dad" {
		t.Fatal(r)
	}
}
