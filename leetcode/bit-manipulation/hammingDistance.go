/* https://leetcode.com/problems/hamming-distance/#/description
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
    0 ≤ x, y < 2^31.

Example:

    Input: x = 1, y = 4
    Output: 2

    Explanation:
        1   (0 0 0 1)
        4   (0 1 0 0)
               ↑   ↑
        The above arrows point to positions where the corresponding bits are different.
在信息论中，两个等长字符串之间的汉明距离（英语：Hamming distance）是两个字符串对应位置的不同字符的个数。
换句话说，它就是将一个字符串变换成另外一个字符串所需要替换的字符个数。
*/

package leetcode

func hammingDistance(x int, y int) int {
	xl, yl := make([]int, 32, 32), make([]int, 32, 32)
	for i := 0; x > 0 || y > 0; i++ {
		xl[i], yl[i] = x%2, y%2
		x, y = x>>1, y>>1
	}

	r := 0
	for i := 0; i < 32; i++ {
		if xl[i] != yl[i] {
			r++
		}
	}
	return r
}
