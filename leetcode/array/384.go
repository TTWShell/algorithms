/* https://leetcode.com/problems/shuffle-an-array/
Shuffle a set of numbers without duplicates.

Example:

// Init an array with set 1, 2, and 3.
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
solution.shuffle();

// Resets the array back to its original configuration [1,2,3].
solution.reset();

// Returns the random shuffling of array [1,2,3].
solution.shuffle();
*/

package larray

import "math/rand"

type Solution struct {
	tmp    []int
	origin []int
}

func Constructor(nums []int) Solution {
	solution := Solution{origin: nums}
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	solution.tmp = tmp
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.tmp), func(i, j int) {
		this.tmp[i], this.tmp[j] = this.tmp[j], this.tmp[i]
	})
	return this.tmp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
