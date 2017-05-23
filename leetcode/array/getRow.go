/* https://leetcode.com/problems/pascals-triangle-ii/#/description
Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].

Note:
Could you optimize your algorithm to use only O(k) extra space?
*/

package leetcode

func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)

	if rowIndex >= 0 {
		r[0] = 1
	}

	for row := 1; row <= rowIndex; row++ {
		r[0], r[row] = 1, 1
		temp := r[0]
		for i := 1; i < row; i++ {
			t := temp + r[i]
			temp, r[i] = r[i], t
		}
	}
	return r
}
