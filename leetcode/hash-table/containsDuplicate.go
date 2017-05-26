/* https://leetcode.com/problems/contains-duplicate/#/solutions
Given an array of integers, find if the array contains any duplicates.
Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.
*/

package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

/*
另外一种解法（时间和上面的一致，都是最高）
func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for index, value := range nums {
        if index > 0 && value == nums[index-1] {
            return true
        }
    }
    return false
}
*/
