package lbm

import "testing"

func Test_getSum(t *testing.T) {
	as := []int{1, 2, 0, 3, 4, 5, 67, 789, 1234}
	bs := []int{2, 3, 4, 5, 6, 7, 88, 123, 5678}

	for i := range as {
		if r := getSum(as[i], bs[i]); r != as[i]+bs[i] {
			t.Fatalf("%d + %d = %d", as[i], bs[i], r)
		}
	}
}
