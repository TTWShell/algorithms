/* https://leetcode.com/problems/single-number/#/description
Given an array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

https://discuss.leetcode.com/topic/22068/easy-java-solution-tell-you-why-using-bitwise-xor
We use bitwise XOR to solve this problem :

first , we have to know the bitwise XOR:

    0 ^ N = N
    N ^ N = 0

So..... if N is the single number

	N1 ^ N1 ^ N2 ^ N2 ^..............^ Nx ^ Nx ^ N
	= (N1^N1) ^ (N2^N2) ^..............^ (Nx^Nx) ^ N
	= 0 ^ 0 ^ ..........^ 0 ^ N
	= N
*/

package leetcode

func singleNumber(nums []int) int {
	/*	maps := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			maps[nums[i]] += 1
		}
		for k, v := range maps {
			if v == 1 {
				return k
			}
		}
		panic("input is error: not every element appears twice except for one")
	*/
	r := 0

	for i := 0; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}
