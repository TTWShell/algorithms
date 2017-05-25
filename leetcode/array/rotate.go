/* https://leetcode.com/problems/rotate-array/#/description
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.

主要是换位置，以下三种思路：

解法一：

1 2 3 4 5 6 7
1 2 3 1 5 6 7
1 2 3 1 5 6 4
1 2 7 1 5 6 4
1 2 7 1 5 3 4
1 6 7 1 5 3 4
1 6 7 1 2 3 4
5 6 7 1 2 3 4

解法二：
先把前n-k个数字翻转一下，再把后k个数字翻转一下，最后再把整个数组翻转。
1 2 3 4 5 6 7
4 3 2 1 5 6 7
4 3 2 1 7 6 5
5 6 7 1 2 3 4

解法三：
通过不停的交换某两个数字的位置来实现旋转
1 2 3 4 5 6 7
5 2 3 4 1 6 7
5 6 3 4 1 2 7
5 6 7 4 1 2 3
5 6 7 1 4 2 3
5 6 7 1 2 4 3
5 6 7 1 2 3 4
*/

package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	c := make([]int, n)
	for i, _ := range nums {
		c[i] = nums[i]
	}

	for i, v := range c {
		nums[(i+k)%n] = v
	}
}
