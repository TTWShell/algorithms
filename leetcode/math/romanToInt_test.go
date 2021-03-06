package lmath

import "testing"

func Test_romanToInt(t *testing.T) {
	nums := []string{"DCXXI"}
	results := []int{621}
	for index, num := range nums {
		if r := romanToInt(num); r != results[index] {
			t.Fatal(num, r, results[index])
		}
	}
}
