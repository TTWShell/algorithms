package search

import "testing"

func Test_BinarySearch(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, target := range sorted {
		if result := BinarySearch(sorted, target); index != result {
			t.Logf("error: target is number `%d`, index is %d, result is `%d`.\n", target, index, result)
			t.Fail()
		}
	}

	if r := BinarySearch(sorted, 10); r != -1 {
		t.Log("error when search target not in sortedData")
		t.Fail()
	}
	if r := BinarySearch(sorted, 0); r != -1 {
		t.Log("error when search target not in sortedData")
		t.Fail()
	}
}
