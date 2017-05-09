package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	inputs := []string{"[", "()", "{()}[]", "{[](}"}
	results := []bool{false, true, true, false}
	for i := 0; i < len(inputs); i++ {
		if r := isValid(inputs[i]); r != results[i] {
			t.Log(inputs[i], results[i], r)
			t.Fail()
		}
	}
}
