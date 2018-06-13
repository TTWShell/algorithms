package lstring

import "testing"

func Test_reverseWords(t *testing.T) {
	if r := reverseWords("Let's take LeetCode contest"); r != "s'teL ekat edoCteeL tsetnoc" {
		t.Fatal(r)
	}
}
