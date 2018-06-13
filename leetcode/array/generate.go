/* https://leetcode.com/problems/pascals-triangle/#/description
Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return

[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package larray

func generate(numRows int) [][]int {
	r := [][]int{}
	if numRows >= 1 {
		r = append(r, []int{1})
	}

	for row := 1; row < numRows; row++ {
		temp := make([]int, len(r[row-1])+1)
		temp[0], temp[len(temp)-1] = 1, 1

		for i := 1; i < len(temp)-1; i++ {
			temp[i] = r[row-1][i-1] + r[row-1][i]
		}
		r = append(r, temp)
	}
	return r
}
