package leetcode

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	if r := myPow(34.00515, -3); math.Abs(r-0.00003) > 0.00001 {
		t.Fatal(r)
	}

	if r := myPow(8.88023, 3); math.Abs(r-700.28148) > 0.00001 {
		t.Fatal(r)
	}
}
