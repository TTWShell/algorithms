package leetcode

import "testing"

func Test_replaceWords(t *testing.T) {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	output := "the cat was rat by the bat"

	if r := replaceWords(dict, sentence); r != output {
		t.Fatal(r)
	}
}
