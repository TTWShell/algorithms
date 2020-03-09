/* https://leetcode.com/problems/candy/
There are N children standing in a line.
Each child is assigned a rating value.

You are giving candies to these children subjected to the
following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

Example 1:

	Input: [1,0,2]
	Output: 5
	Explanation: You can allocate to the first,
	second and third child with 2, 1, 2 candies respectively.

Example 2:

	Input: [1,2,2]
	Output: 4
	Explanation: You can allocate to the first, second and third child
	with 1, 2, 1 candies respectively.
	The third child gets 1 candy because it satisfies the above two
	conditions.
*/

package lgreedy

func candy(a []int) int {
	// 头尾各处理一次，就不需要关注递减区间的特殊处理
	length := len(a)
	if length < 2 {
		return length
	}

	candies := make([]int, length, length)
	candies[0] = 1
	for i := 1; i < length; i++ {
		if a[i] > a[i-1] {
			candies[i] = 1 + candies[i-1]
		} else {
			candies[i] = 1
		}
	}

	for i := length - 2; i >= 0; i-- {
		if a[i] > a[i+1] && candies[i] <= candies[i+1] {
			candies[i] = 1 + candies[i+1]
		}
	}

	res := 0
	for i := 0; i < length; i++ {
		res += candies[i]
	}
	return res
}
