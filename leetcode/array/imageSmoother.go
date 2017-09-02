/* https://leetcode.com/problems/image-smoother/description/
Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

Example 1:
    Input:
    [[1,1,1],
     [1,0,1],
     [1,1,1]]
    Output:
    [[0, 0, 0],
     [0, 0, 0],
     [0, 0, 0]]
    Explanation:
    For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
    For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
    For the point (1,1): floor(8/9) = floor(0.88888889) = 0
Note:
    The value in the given matrix is in the range of [0, 255].
    The length and width of the given matrix are in the range of [1, 150].
*/

package leetcode

func imageSmoother(M [][]int) [][]int {
	colLen := len(M)
	rowLen := len(M[0])
	res := make([][]int, colLen)
	for i := range res {
		res[i] = make([]int, rowLen, rowLen)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for row := range M {
		for col := range M[row] {
			sum, count := 0, 0
			for i := max(0, row-1); i < min(colLen, row+2); i++ {
				for j := max(0, col-1); j < min(rowLen, col+2); j++ {
					sum += M[i][j]
					count++
				}
			}
			res[row][col] = int(sum / count)
		}
	}
	return res
}
