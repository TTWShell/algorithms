/* https://leetcode.com/problems/contains-duplicate-iii/description/
Given an array of integers, find out whether there are two distinct indices i and j in the array such that the
absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.
*/

package leetcode

func containsNearbyAlmostDuplicate(v []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}

	bucket := func(x int, t int) int {
		if x >= 0 {
			return x / t
		}
		return x/t - 1
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	check := func(b map[int]int, x int, t int, nb bool) bool {
		p := bucket(x, t)
		if _, ok := b[p]; ok {
			return true
		}
		if !nb {
			return false
		}
		if l, ok := b[p-1]; ok && abs(l-x) <= t {
			return true
		}
		if r, ok := b[p+1]; ok && abs(r-x) <= t {
			return true
		}
		return false
	}

	add := func(b map[int]int, x int, t int) {
		b[bucket(x, t)] = x
	}

	remove := func(b map[int]int, x int, t int) {
		delete(b, bucket(x, t))
	}

	nb := true
	if t == 0 {
		t = 1
		nb = false
	}

	b := map[int]int{}
	for i := 0; i <= k && i < len(v); i++ {
		if check(b, v[i], t, nb) {
			return true
		}
		add(b, v[i], t)
	}
	for i := k + 1; i < len(v); i++ {
		remove(b, v[i-k-1], t)
		if check(b, v[i], t, nb) {
			return true
		}
		add(b, v[i], t)
	}
	return false
}
