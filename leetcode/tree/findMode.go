/* https://leetcode.com/problems/find-mode-in-binary-search-tree/#/description
Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

    The left subtree of a node contains only nodes with keys less than or equal to the node's key.
    The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
    Both the left and right subtrees must also be binary search trees.

For example:
    Given BST [1,null,2,2],
       1
        \
         2
        /
       2
    return [2].

Note: If a tree has more than one mode, you can return them in any order.

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
*/

package leetcode

func findMode(root *TreeNode) []int {
	var (
		maxCount, currentVal, tempCount int
		result                          []int
	)
	return inorder(root, &maxCount, &currentVal, &tempCount, result)
}

func inorder(root *TreeNode, maxCount, currentVal, tempCount *int, result []int) []int {
	if root == nil {
		return result
	}
	result = inorder(root.Left, maxCount, currentVal, tempCount, result)
	if root.Val != *currentVal {
		*currentVal = root.Val
		*tempCount = 0
	}
	*tempCount++
	if *tempCount > *maxCount {
		*maxCount = *tempCount
		result = []int{root.Val}
	} else if *tempCount == *maxCount {
		result = append(result, root.Val)
	}
	return inorder(root.Right, maxCount, currentVal, tempCount, result)
}
