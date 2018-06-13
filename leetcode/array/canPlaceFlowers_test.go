package larray

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	if r := canPlaceFlowers(flowerbed, 1); r != true {
		t.Fatal(r)
	}

	if r := canPlaceFlowers(flowerbed, 2); r != false {
		t.Fatal(r)
	}

	if r := canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2); r != true {
		t.Fatal(r)
	}
}
