package lmath

import "testing"

func Test_intToRoman(t *testing.T) {
	output := []string{"DCXXI"}
	input := []int{621}
	for index, num := range input {
		if r := intToRoman(num); r != output[index] {
			t.Fatal(num, r, output[index])
		}
	}
}
