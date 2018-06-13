package larray

import (
	"math"
	"testing"
)

func Test_findMaxAverage2(t *testing.T) {
	input := []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 4); math.Abs(12.75-r) > 0.00001 {
		t.Fatal(r)
	}

	input = []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 5); math.Abs(10.8-r) > 0.00001 {
		t.Fatal(r)
	}

	input = []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 6); math.Abs(9.166666666666666-r) > 0.00001 {
		t.Fatal(r)
	}

	input = []int{8, 9, 3, 1, 8, 3, 0, 6, 9, 2}
	if r := findMaxAverage2(input, 8); math.Abs(5.222222222222222-r) > 0.00001 {
		t.Fatal(r)
	}
}
