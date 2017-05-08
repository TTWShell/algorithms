package stringmatching

import "testing"

func Test_Naive(t *testing.T) {
	sList := []string{"BBC ABCDAB ABCDABCDABDE", "", "ab", "abc"}
	txtList := []string{"ABCDABD", "a", "c", "c"}
	result := []bool{true, false, false, true}
	for i := 0; i < len(sList); i++ {
		if Naive(sList[i], txtList[i]) != result[i] {
			t.Log(sList[i], txtList[i], result[i])
			t.Fail()
		}
	}
}
