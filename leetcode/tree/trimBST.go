/* https://leetcode.com/problems/trim-a-binary-search-tree/description/
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L).
You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
    Input:
        1
       / \
      0   2

      L = 1
      R = 2

    Output:
        1
          \
           2
Example 2:
    Input:
        3
       / \
      0   4
       \
        2
       /
      1

      L = 1
      R = 3

    Output:
          3
         /
       2
      /
     1
*/

package ltree

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	// https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%85%83%E6%90%9C%E5%B0%8B%E6%A8%B9
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left, root.Right = trimBST(root.Left, L, R), trimBST(root.Right, L, R)
	return root
}
