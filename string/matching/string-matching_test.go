package stringmatching

import "testing"

func Test_Naive(t *testing.T) {
	sList := []string{"BBC ABCDAB ABCDABCDABDE", "", "ab", "abc"}
	txtList := []string{"ABCDABD", "a", "c", "c"}
	resultList := []bool{true, false, false, true}
	for i := 0; i < len(sList); i++ {
		if result := Naive(sList[i], txtList[i]); result != resultList[i] {
			t.Fatal(sList[i], txtList[i], resultList[i], "result is:", result)
		}
	}
}

func Test_KnuthMorrisPratt(t *testing.T) {
	sList := []string{"BBC ABCDAB ABCDABCDABDE", "", "ab", "abc"}
	txtList := []string{"ABCDABD", "a", "c", "c"}
	resultList := []bool{true, false, false, true}
	for i := 0; i < len(sList); i++ {
		if result := KnuthMorrisPratt(sList[i], txtList[i]); result != resultList[i] {
			t.Fatal(sList[i], txtList[i], resultList[i], "result is:", result)
		}
	}
}
