/* https://leetcode.com/problems/subtree-of-another-tree/#/description
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s.
A subtree of s is a tree consists of a node in s and all of this node's descendants.
The tree s could also be considered as a subtree of itself.

Example 1:
    Given tree s:

         3
        / \
       4   5
      / \
     1   2
    Given tree t:
       4
      / \
     1   2
    Return true, because t has the same structure and node values with a subtree of s.

Example 2:
    Given tree s:

         3
        / \
       4   5
      / \
     1   2
        /
       0
    Given tree t:
       4
      / \
     1   2
    Return false.
*/

package ltree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var isSameTree func(p *TreeNode, q *TreeNode) bool
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		return (p == nil && q == nil) ||
			(p != nil && q != nil &&
				p.Val == q.Val &&
				isSameTree(p.Left, q.Left) &&
				isSameTree(p.Right, q.Right))
	}

	if isSameTree(s, t) ||
		s.Left != nil && isSameTree(s.Left, t) ||
		s.Right != nil && isSameTree(s.Right, t) {
		return true
	}

	return (s.Left != nil && isSubtree(s.Left, t)) || (s.Right != nil && isSubtree(s.Right, t))
}
