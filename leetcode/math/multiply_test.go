package lmath

import "testing"

func Test_multiply(t *testing.T) {
	if r := multiply("12141414", "6563457472"); r != "79689654438945408" {
		t.Fatal(r)
	}
}
