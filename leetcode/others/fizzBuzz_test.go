package leetcode

import "testing"

func Test_fizzBuzz(t *testing.T) {
	result := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	r := fizzBuzz(15)

	if len(result) != len(r) {
		t.Fatal(r)
	}
	for i := range result {
		if result[i] != r[i] {
			t.Fatal(r, r[i], result[i])
		}
	}
}
