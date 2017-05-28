/* https://leetcode.com/problems/binary-tree-paths/#/description
Given a binary tree, return all root-to-leaf paths.

For example, given the following binary tree:

   1
 /   \
2     3
 \
  5
All root-to-leaf paths are:

["1->2->5", "1->3"]
*/

package leetcode

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	r := []string{}
	if root != nil {
		if root.Left != nil {
			for _, s := range binaryTreePaths(root.Left) {
				r = append(r, strconv.Itoa(root.Val)+"->"+s)
			}
		}
		if root.Right != nil {
			for _, s := range binaryTreePaths(root.Right) {
				r = append(r, strconv.Itoa(root.Val)+"->"+s)
			}
		}
		if root.Left == nil && root.Right == nil {
			r = append(r, strconv.Itoa(root.Val))
		}
	}
	return r
}
