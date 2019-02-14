/* https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]

https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png

Example 1:

	Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	Output: 6
	Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:

	Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
	Output: 2
	Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Note:

	All of the nodes' values will be unique.
	p and q are different and both values will exist in the BST.
*/

package ltree

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	target := [][]int{}
	cache := [][]*TreeNode{[]*TreeNode{root}}
	flag := true
	for tmp := 0; flag == true; tmp++ {
		flag = false
		cur := []*TreeNode{}
		for idx, node := range cache[tmp] {
			if node == nil {
				cur = append(cur, nil)
				cur = append(cur, nil)
			} else {
				if node.Val == p.Val || node.Val == q.Val {
					target = append(target, []int{tmp, idx})
					if len(target) == 2 {
						break
					}
				}
				if node.Left != nil || node.Right != nil {
					flag = true
				}
				cur = append(cur, node.Left)
				cur = append(cur, node.Right)
			}
		}
		cache = append(cache, cur)
	}

	x1, y1 := target[0][0], target[0][1]
	x2, y2 := target[1][0], target[1][1]
	for x1 != x2 || y1 != y2 {
		if x1 != x2 {
			if x1 > x2 {
				x1, y1 = x1-1, y1/2
			} else {
				x2, y2 = x2-1, y2/2
			}
		} else {
			if y1 != y2 {
				x1, y1 = x1-1, y1/2
				x2, y2 = x2-1, y2/2
			}
		}
	}
	return cache[x1][y1]
}
