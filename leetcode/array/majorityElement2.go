/* https://leetcode.com/problems/majority-element-ii/description/
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times. The algorithm should run in linear time and in O(1) space.
*/

package leetcode

func majorityElement2(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			candidate2, count2 = num, 1
		} else {
			count1--
			count2--
		}
	}

	res, c1, c2 := []int{}, 0, 0
	for _, num := range nums {
		if num == candidate1 {
			c1++
		} else if num == candidate2 {
			c2++
		}
	}
	if c1 > len(nums)/3 {
		res = append(res, candidate1)
	}
	if c2 > len(nums)/3 {
		res = append(res, candidate2)
	}
	return res
}
