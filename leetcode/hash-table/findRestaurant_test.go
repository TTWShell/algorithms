package leetcode

import "testing"

func Test_findRestaurant(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	if r := findRestaurant(list1, list2); len(r) != 1 || r[0] != "Shogun" {
		t.Fatal(r)
	}

	list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 = []string{"KFC", "Shogun", "Burger King"}
	if r := findRestaurant(list1, list2); len(r) != 1 || r[0] != "Shogun" {
		t.Fatal(r)
	}
}
